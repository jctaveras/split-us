package userhandler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jctaveras/split-us/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileHandler interface {
	Handler(primitive.ObjectID) http.HandlerFunc
}

type profileHandler struct {}

func NewProfileHandler() ProfileHandler {
	return &profileHandler{}
}

func (handler *profileHandler) Handler(id primitive.ObjectID) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var user User
		storage := req.Context().Value(database.Storage{}).(*database.Storage)

		if error := storage.FindUser(bson.D{{Key: "_id", Value: id}}).Decode(&user); error == mongo.ErrNoDocuments {
			http.Error(res, error.Error(), http.StatusNotFound)
			return
		} else if error != nil {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}

		if data, error := json.Marshal(user); error != nil && error != io.EOF {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		} else {
			res.Header().Add("Content-Type", "application/json")
			res.Write(data)
		}
	}
}

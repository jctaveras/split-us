package servicehandler

import (
	"encoding/json"
	"net/http"

	"github.com/jctaveras/split-us/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddUserHandler interface {
	Handler(primitive.ObjectID) http.HandlerFunc
}

type addUserHandler struct {}

func NewAddUserHandler() AddUserHandler {
	return &addUserHandler{}
}

func (handler *addUserHandler) Handler(id primitive.ObjectID) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		storage := req.Context().Value(database.Storage{}).(*database.Storage)
		var service Service
		var userIds []primitive.ObjectID

		if error := storage.FindService(bson.D{{Key: "_id", Value: id}}).Decode(&service); error == mongo.ErrNoDocuments {
			http.Error(res, "Service Not Found", http.StatusNotFound)
			return
		} else if error != nil {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}

		if error := json.NewDecoder(req.Body).Decode(&userIds); error != nil {
			http.Error(res, error.Error(), http.StatusBadRequest)
			return
		}

		for _, userId := range userIds {
			var user User
			if error := storage.FindUser(bson.D{{Key: "_id", Value: userId}}).Decode(&user); error == mongo.ErrNoDocuments {
				http.Error(res, "User Not Found with ID: " + userId.Hex(), http.StatusNotFound)
				return
			} else if error != nil {
				http.Error(res, error.Error(), http.StatusBadRequest)
				return
			}

			if error := storage.AddUserToService(service.ID, user); error != nil {
				http.Error(res, error.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

package servicehandler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jctaveras/split-us/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddServiceHandler interface {
	Handler(primitive.ObjectID) http.HandlerFunc
}

type addServiceHandler struct{}

func NewAddServiceHandler() AddServiceHandler {
	return &addServiceHandler{}
}

func (handler *addServiceHandler) Handler(id primitive.ObjectID) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		storage := req.Context().Value(database.Storage{}).(*database.Storage)
		var user User
		
		if error := storage.FindUser(bson.D{{Key: "_id", Value: id}}).Decode(&user); error == mongo.ErrNoDocuments {
			http.Error(res, "No User Found with ID: " + id.Hex(), http.StatusNotFound)
			return
		} else if error != nil {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}

		var service Service
		
		if error := json.NewDecoder(req.Body).Decode(&service); error != nil {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}

		service.CreatedAt = time.Now()
		service.CreatedBy = user
		service.UpdatedAt = time.Now()

		if error := validator.New().Struct(service); error != nil {
			http.Error(res, error.Error(), http.StatusBadRequest)
			return
		}

		if error := storage.NewService(service); error != nil {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}

		res.WriteHeader(http.StatusCreated)
	}
}

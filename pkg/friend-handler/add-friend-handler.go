package friendhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jctaveras/split-us/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddFriendHandler interface {
	Handler(primitive.ObjectID) http.HandlerFunc 
}

type addFriendHandler struct {}

func NewAddFriendHandler() AddFriendHandler {
	return &addFriendHandler{}
}

func (handler *addFriendHandler) Handler(id primitive.ObjectID) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var data struct{Friends []string `json:"friends"`}

		if error := json.NewDecoder(req.Body).Decode(&data); error != nil {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}

		var user User
		storage := req.Context().Value(database.Storage{}).(*database.Storage)

		if error := storage.FindUser(bson.D{{Key: "_id", Value: id}}).Decode(&user); error == mongo.ErrNoDocuments {
			http.Error(res, fmt.Sprintf("No user Found with ID: %v", id), http.StatusNotFound)
			return
		} else if error != nil {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}
		
		for _, friendId := range data.Friends {
			var friend User
			friendObjectId, error := primitive.ObjectIDFromHex(friendId)

			if error != nil {
				http.Error(res, "Invalida Friend ID", http.StatusBadRequest)
				return
			}

			if error := storage.FindUser(bson.D{{Key: "_id", Value: friendObjectId}}).Decode(&friend); error == mongo.ErrNoDocuments {
				http.Error(res, fmt.Sprintf("No User Found with ID: %v", friendObjectId), http.StatusNotFound)
				return
			} else if error != nil {
				http.Error(res, error.Error(), http.StatusInternalServerError)
				return	
			}

			if error := storage.AddFriend(id, friend); error != nil {
				http.Error(res, error.Error(), http.StatusInternalServerError)
				return
			}

			if error := storage.AddFriend(friend.ID, user); error != nil {
				http.Error(res, error.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

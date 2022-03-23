package friendhandler

import (
	"encoding/json"
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

		storage := req.Context().Value(database.Storage{}).(*database.Storage)
		
		for _, friendId := range data.Friends {
			var friend User
			friendObjectId, error := primitive.ObjectIDFromHex(friendId)

			if error != nil {
				http.Error(res, error.Error(), http.StatusInternalServerError)
				return
			}

			if error := storage.FindUser(bson.D{{Key: "_id", Value: friendObjectId}}).Decode(&friend); error == mongo.ErrNoDocuments {
				http.Error(res, error.Error(), http.StatusNotFound)
				return
			} else if error != nil {
				http.Error(res, error.Error(), http.StatusInternalServerError)
				return	
			}

			if error := storage.AddFriend(id, friend); error != nil {
				http.Error(res, error.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

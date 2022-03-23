package friendhandler

import (
	"net/http"

	"github.com/jctaveras/split-us/internal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const AddFriend = "add-friends"

type FriendHandler struct {
	AddFriendHandler AddFriendHandler
}

func (handler *FriendHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = internal.ShiftPath(req.URL.Path)
	id, error := primitive.ObjectIDFromHex(head)

	if error != nil {
		http.Error(res, "Invalid User ID", http.StatusBadRequest)
		return
	}

	if req.URL.Path != "/" {
		head, req.URL.Path = internal.ShiftPath(req.URL.Path)

		switch req.Method {
		case http.MethodPatch:
			switch head {
			case AddFriend:
				handler.AddFriendHandler.Handler(id).ServeHTTP(res, req)
			default:
				http.Error(res, "Not Found", http.StatusNotFound)
			}
		default:
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
		
		return
	}
}

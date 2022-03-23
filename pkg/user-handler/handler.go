package userhandler

import (
	"net/http"

	"github.com/jctaveras/split-us/internal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	UserProfile = "profile"
)

type UserHandler struct {
	ProfileHandler ProfileHandler
}

func (handler *UserHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var hexID, head string
	hexID, req.URL.Path = internal.ShiftPath(req.URL.Path)
	id, error := primitive.ObjectIDFromHex(hexID)

	if error != nil {
		http.Error(res, "Invalid User ID", http.StatusBadRequest)
		return
	}

	if req.URL.Path != "/" {
		head, req.URL.Path = internal.ShiftPath(req.URL.Path)

		switch req.Method {
		case http.MethodGet:
			switch head {
			case UserProfile:
				handler.ProfileHandler.Handler(id).ServeHTTP(res, req)
			default:
				http.Error(res, "Not Found", http.StatusNotFound)
			}
		default:
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		}

		return
	}
}

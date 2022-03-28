package servicehandler

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/jctaveras/split-us/internal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	AddService = "add-service"
	AddUser = "add-user"
)

type ServiceHandler struct {
	AddServiceHandler AddServiceHandler
	AddUserHandler AddUserHandler
}

type AuthTokenClaim struct {
	UserID string `json:"userID"`
	Exp	string `json:"exp"`
	jwt.StandardClaims
}

func (handler *ServiceHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	authToken := req.Header.Get("Authorization")
	token, error := jwt.ParseWithClaims(authToken, &AuthTokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	claim, ok := token.Claims.(*AuthTokenClaim)

	if !(ok && token.Valid) {
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return 
	}

	id, error := primitive.ObjectIDFromHex(claim.UserID)

	if error != nil {
		http.Error(res, "Invalid User ID", http.StatusUnauthorized)
		return
	}

	var head string
	head, req.URL.Path = internal.ShiftPath(req.URL.Path)

	if req.URL.Path != "/" {
		serviceId, error := primitive.ObjectIDFromHex(head)

		if error != nil {
			http.Error(res, "Invalid Service ID", http.StatusBadRequest)
			return
		}

		head, req.URL.Path = internal.ShiftPath(req.URL.Path)
		
		switch req.Method {
		case http.MethodPatch:
			switch head {
			case AddUser:
				handler.AddUserHandler.Handler(serviceId).ServeHTTP(res, req)
			default:
				http.Error(res, "Not Found", http.StatusNotFound)
			}
		default:
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		}

		return
	}
	
	switch req.Method {
	case http.MethodPost:
		switch head {
		case AddService:
			handler.AddServiceHandler.Handler(id).ServeHTTP(res, req)
		default:
			http.Error(res, "Not Found", http.StatusNotFound)
		}
	default:
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

}

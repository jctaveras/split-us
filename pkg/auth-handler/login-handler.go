package authhandler

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jctaveras/split-us/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler interface {
	Handler() http.HandlerFunc
}

type loginHandler struct {}

func NewLoginHandler() LoginHandler {
	return &loginHandler{}
}

func (handler *loginHandler) Handler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var credentials User
		var user User

		if error := json.NewDecoder(req.Body).Decode(&credentials); error != nil {
			http.Error(res, "Invalid Request", http.StatusBadRequest)
			return
		}

		storage := req.Context().Value(database.Storage{}).(*database.Storage)
		
		if error := storage.FindUser(bson.D{{Key: "email", Value: credentials.Email}}).Decode(&user); error == mongo.ErrNoDocuments {
			http.Error(res, error.Error(), http.StatusNotFound)
			return
		} else if error != nil {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}

		if didPasswordMatch([]byte(user.Password), []byte(credentials.Password)) {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"userID": user.ID.Hex(),
				"exp": time.Now().Add(5 * time.Minute),
			})
			tokenString, signedError := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

			if signedError != nil {
				http.Error(res, signedError.Error(), http.StatusInternalServerError)
				return
			}

			res.Write([]byte(tokenString))
			return
		}

		http.Error(res, "Invalida Password", http.StatusBadRequest)
	}
}

func didPasswordMatch(hPwd []byte, pwd []byte) bool {
	return bcrypt.CompareHashAndPassword(hPwd, pwd) == nil
}

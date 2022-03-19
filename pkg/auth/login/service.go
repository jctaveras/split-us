package login

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Storage interface {
	FindUser(interface{}) *mongo.SingleResult
}

type service struct {
	s Storage
}

func NewLoginService(s Storage) *service {
	return &service{s}
}

func (service *service) Login(credentials User) (string, error) {
	result := service.s.FindUser(bson.D{{Key: "email", Value: credentials.Email}})
	var user User

	if error := result.Decode(&user); error != nil {
		return "", error
	} else {
		if didPasswordMatch([]byte(user.Password), []byte(credentials.Password)) {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"userID": user.ID.Hex(),
				"exp": time.Now().Add(5 * time.Minute),
			})
			tokenString, signedError := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

			if signedError != nil {
				return "", signedError
			}

			return tokenString, nil
		}

		return "", errors.New("Invalid Password.")
	}
}

func didPasswordMatch(hPwd []byte, pwd []byte) bool {
	return bcrypt.CompareHashAndPassword(hPwd, pwd) == nil
}

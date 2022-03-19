package login

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Storage interface {
	FindUser(User) (User, error)
}

type service struct {
	s Storage
}

func NewLoginService(s Storage) *service {
	return &service{s}
}

func (service *service) Login(credentials User) (string, error) {
	if user, error := service.s.FindUser(credentials); error != nil {
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

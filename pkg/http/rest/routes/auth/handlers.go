package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jctaveras/split-us/pkg/database/signup"
	"github.com/jctaveras/split-us/pkg/database/storage"
	"github.com/jctaveras/split-us/pkg/http/router"
	"golang.org/x/crypto/bcrypt"
)

func InitAuthHandlers(ctx context.Context) {
	router.Routes.POST("/api/user/sign-up", func(w http.ResponseWriter, r *http.Request) {
		storage := ctx.Value(storage.Storage{}).(signup.Storage)
		service := signup.NewSignUpService(storage)
		hashChan := make(chan []byte)
		var userData signup.User

		error := json.NewDecoder(r.Body).Decode(&userData)

		if error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}

		// This is just to understand Channels 
		go func (password string, ch chan []byte)  {
			if hash, error := bcrypt.GenerateFromPassword([]byte(password), 10); error != nil {
				log.Fatalf("Error while hashing the password: %v", error)
			} else {
				ch <- hash
			}
		}(userData.Password, hashChan)
			
		userData.Password = string(<-hashChan)

		if error := validator.New().Struct(userData); error != nil {
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}

		if error := service.SignUp(userData); error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
}

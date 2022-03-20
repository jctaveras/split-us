package auth

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jctaveras/split-us/pkg/auth/login"
	"github.com/jctaveras/split-us/pkg/auth/signup"
	"github.com/jctaveras/split-us/pkg/database"
	"github.com/jctaveras/split-us/pkg/http/router"
)

func InitAuthHandlers(ctx context.Context) {
	router.Routes.POST("/api/user/sign-up", func(w http.ResponseWriter, r *http.Request) {
		storage := ctx.Value(database.Storage{}).(signup.Storage)
		service := signup.NewSignUpService(storage)
		var userData signup.User

		if error := json.NewDecoder(r.Body).Decode(&userData); error != nil && error != io.EOF {
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}

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

	router.Routes.POST("/api/user/login", func(w http.ResponseWriter, r *http.Request) {
		storage := ctx.Value(database.Storage{}).(login.Storage)
		service := login.NewLoginService(storage)
		var credentials login.User

		if error := json.NewDecoder(r.Body).Decode(&credentials); error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}

		if error := validator.New().Struct(credentials); error != nil {
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}

		token, error := service.Login(credentials)

		if error != nil {
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}

		w.Write([]byte(token))
	})
}

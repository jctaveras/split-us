package authhandler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jctaveras/split-us/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

type SignUpHandler interface {
	Handler() http.HandlerFunc
}

type signUpHandler struct {}

func NewSignUpHandler() SignUpHandler {
	return &signUpHandler{}
}

func (handler *signUpHandler) Handler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var data User

		if error := json.NewDecoder(req.Body).Decode(&data); error != nil && error != io.EOF {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}

		if error := validator.New().Struct(data); error != nil {
			http.Error(res, error.Error(), http.StatusBadRequest)
			return
		}

		storage := req.Context().Value(database.Storage{}).(*database.Storage)
		hashed, error := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

		if error != nil {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}

		data.Password = string(hashed)

 		if error := storage.NewUser(data); error != nil {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}

		res.WriteHeader(http.StatusCreated)
	}
}

package authhandler

import (
	"net/http"

	"github.com/jctaveras/split-us/internal"
)

const (
	SignUp = "sign-up"
	Login  = "login"
)

type AuthHandler struct {
	LoginHandler LoginHandler
	SignUpHandler SignUpHandler
}

func (handler *AuthHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string

	head, req.URL.Path = internal.ShiftPath(req.URL.Path)

	if req.Method == http.MethodPost {
		switch head {
		case Login:
			handler.LoginHandler.Handler().ServeHTTP(res, req)
		case SignUp:
			handler.SignUpHandler.Handler().ServeHTTP(res, req)
		default:
			http.Error(res, "Not Found", http.StatusNotFound)
		}
	}
}

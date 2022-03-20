package user

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/jctaveras/split-us/pkg/database"
	"github.com/jctaveras/split-us/pkg/http/router"
	"github.com/jctaveras/split-us/pkg/user"
)

func InitUserHandlers(ctx context.Context) {
	router.Routes.GET("/api/user/profile", func(w http.ResponseWriter, r *http.Request) {
		storage := ctx.Value(database.Storage{}).(user.Storage)
		userService := user.NewUserService(storage)
		var userData user.User
		
		if error := json.NewDecoder(r.Body).Decode(&userData); error != nil && error != io.EOF {
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}

		if data, error := userService.Profile(userData); error != nil {
			http.Error(w, error.Error(), http.StatusNotFound)
			return
		} else {
			data, error := json.Marshal(data)

			if error != nil && error != io.EOF {
				http.Error(w, error.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Add("Content-Type", "application/json")
			w.Write(data)
		}
	})
}

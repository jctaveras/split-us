package app

import (
	"context"
	"net/http"

	"github.com/jctaveras/split-us/internal"
	authhandler "github.com/jctaveras/split-us/pkg/auth-handler"
	"github.com/jctaveras/split-us/pkg/database"
	friendhandler "github.com/jctaveras/split-us/pkg/friend-handler"
	userhandler "github.com/jctaveras/split-us/pkg/user-handler"
)

type App struct {
	AuthHandler   *authhandler.AuthHandler
	FriendHandler *friendhandler.FriendHandler
	UserHandler   *userhandler.UserHandler
	Repository    *database.Storage
}

func (app *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = internal.ShiftPath(req.URL.Path)
	req = req.WithContext(context.WithValue(req.Context(), database.Storage{}, app.Repository))

	switch head {
	case "auth":
		app.AuthHandler.ServeHTTP(res, req)
	case "friend":
		app.FriendHandler.ServeHTTP(res, req)
	case "user":
		app.UserHandler.ServeHTTP(res, req)
		return
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

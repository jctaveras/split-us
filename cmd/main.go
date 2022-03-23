package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jctaveras/split-us/pkg/app"
	authhandler "github.com/jctaveras/split-us/pkg/auth-handler"
	"github.com/jctaveras/split-us/pkg/database"
	friendhandler "github.com/jctaveras/split-us/pkg/friend-handler"
	userhandler "github.com/jctaveras/split-us/pkg/user-handler"

	"github.com/joho/godotenv"
)

func main() {
	if error := godotenv.Load(); error != nil {
		panic(error)
	}

	s := database.NewStorage()
	app := &app.App{
		AuthHandler: &authhandler.AuthHandler{
			LoginHandler: authhandler.NewLoginHandler(),
			SignUpHandler: authhandler.NewSignUpHandler(),
		},
		FriendHandler: &friendhandler.FriendHandler{
			AddFriendHandler: friendhandler.NewAddFriendHandler(),
		},
		UserHandler: &userhandler.UserHandler{
			ProfileHandler: userhandler.NewProfileHandler(),
		},
		Repository: s,
	}

	defer func() {
		if error := s.Database.Client().Disconnect(context.TODO()); error != nil {
			panic(error)
		}
	}()

	fmt.Println("Server is running on: http://localhost:8080")
	http.ListenAndServe(":8080", app)
}

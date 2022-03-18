package main

import (
	"context"
	"fmt"

	"github.com/jctaveras/split-us/pkg/database/storage"
	"github.com/jctaveras/split-us/pkg/http/rest"
)

type SuccessfullResponse struct {
	Data any `json:"data"`
}

func main() {
	server := rest.NewServer()

	fmt.Println("Server is running on: http://localhost:8080")

	defer func () {
		if error := storage.DataBase.Client().Disconnect(context.TODO()); error != nil {
			panic(error)
		}
	}()

	if error := server.Start(); error != nil {
		panic(error)
	}
}

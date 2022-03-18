package main

import (
	"context"
	"fmt"

	"github.com/jctaveras/split-us/pkg/database/storage"
	"github.com/jctaveras/split-us/pkg/http/rest"
	"github.com/joho/godotenv"
)

func main() {
	if error := godotenv.Load(); error != nil {
		panic(error)
	}

	s := storage.NewStorage()
	ctx := context.WithValue(context.Background(), storage.Storage{}, s)
	server := rest.NewServer()

	fmt.Println("Server is running on: http://localhost:8080")

	defer func () {
		if error := s.Database.Client().Disconnect(context.TODO()); error != nil {
			panic(error)
		}
	}()

	if error := server.Start(ctx); error != nil {
		panic(error)
	}
}

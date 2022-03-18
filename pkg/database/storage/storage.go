package storage

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newStorage() *mongo.Database {
	if error := godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	mongoURI := os.Getenv("MONGO_URI")

	if client, error := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI)); error != nil {
		panic(error)
	} else {
		return client.Database("split-us")
	}
}

var DataBase = newStorage()

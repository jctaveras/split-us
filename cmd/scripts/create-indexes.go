package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if error := godotenv.Load(); error != nil {
		log.Fatal(error)
	}
	
	mongoURI := os.Getenv("MONGO_URI")
	client, error := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))

	if error != nil {
		log.Fatalf("Something went wrong while connecting to the database: %v", error)
	}

	collection := client.Database("split-us").Collection("Users")

	collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{ Key: "email", Value: 1 }},
		Options: options.Index().SetUnique(true),
	})
}

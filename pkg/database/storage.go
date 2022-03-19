package database 

import (
	"context"
	"os"

	"github.com/jctaveras/split-us/pkg/auth/signup"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Database *mongo.Database
}

const (
	UserCollection = "Users"
)

func NewStorage() *Storage {
	mongoURI := os.Getenv("MONGO_URI")

	if client, error := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI)); error != nil {
		panic(error)
	} else {
		return &Storage{Database: client.Database("split-us")}
	}
}

func (s *Storage) SignUp(user signup.User) error {
	collection := s.Database.Collection(UserCollection)

	if _, error := collection.InsertOne(context.TODO(), user); error != nil {
		return error
	}

	return nil
}

func (s *Storage) FindUser(filter interface{}) *mongo.SingleResult {
	collection := s.Database.Collection(UserCollection)

	return collection.FindOne(context.TODO(), filter)
}

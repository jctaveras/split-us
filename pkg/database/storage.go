package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *Storage) NewUser(user interface{}) error {
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

func (s *Storage) AddFriend(id primitive.ObjectID, user interface{}) error {
	collection := s.Database.Collection(UserCollection)
	update := bson.D{{
		Key: "$push", 
		Value: bson.D{{
			Key: "friends", 
			Value: user,
		}},
	}}

	if _, error := collection.UpdateByID(context.TODO(), id, update); error != nil {
		return error
	}

	return nil
}

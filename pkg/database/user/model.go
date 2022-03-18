package user

import (
	"context"

	"github.com/jctaveras/split-us/pkg/database/storage"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model interface {
	NewUser(Schema) error
}

type model struct {
	collection *mongo.Collection
}

func NewModel() Model {
	return &model{
		collection: storage.DataBase.Collection("Users"),
	}
}

func (s *model) NewUser(user Schema) error {
	if _, error := s.collection.InsertOne(context.Background(), user); error != nil {
		return error
	}

	return nil
}

package user

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage interface {
	FindUser(interface{}) *mongo.SingleResult
}

type service struct {
	s Storage
}

func NewUserService(s Storage) *service {
	return &service{s}
}

func (service *service) Profile(user User) (User, error) {
	result := service.s.FindUser(bson.D{{Key: "email", Value: user.Email}})
	var profile User

	if error := result.Decode(&profile); error != nil {
		return profile, error
	}

	return profile, nil
}

package database 

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	DoB       time.Time          `bson:"dob,omitempty" json:"dob,omitempty"`
	Email     string             `bson:"email" json:"email" validate:"required,email"`
	FirstName string             `bson:"firstName" json:"firstName" validate:"required"`
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	LastName  string             `bson:"lastName" json:"lastName" validate:"required"`
	Password  string             `bson:"password" json:"password" validate:"required"`
}

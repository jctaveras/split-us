package login 

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Email     string             `bson:"email" json:"email" validate:"required,email"`
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Password  string             `bson:"password" json:"password" validate:"required"`
}

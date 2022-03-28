package servicehandler

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	FirstName string             `bson:"firstName,omitempty" json:"firstName,omitempty"`
	Friends   []User             `bson:"friends,omitempty" json:"friends,omitempty"`
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	LastName  string             `bson:"lastName,omitempty" json:"lastName,omitempty"`
}

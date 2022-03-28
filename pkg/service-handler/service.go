package servicehandler

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty" validate:"required"`
	CreatedBy User               `bson:"createdBy" json:"createdBy" validate:"required"`
	DeletedAt time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name" validate:"required"`
	Price     float32            `bson:"price" json:"price" validate:"required"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Users     []User             `bson:"users,omitempty" json:"users,omitempty"`
}

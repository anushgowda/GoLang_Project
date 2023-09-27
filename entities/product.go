package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID       primitive.ObjectID      `json:"id" bson:"_id"`
	Name            string           `json:"Name" bson:"Name,required"`
	Category        string           `json:"Category" bson:"Category,required"`
	Quantity        int			     `json:"Quantity" bson:"Quantity,required"`
	CreatedAt     time.Time 	     `json:"CreatedAt" bson:"created_at"`
	UpdatedAt     time.Time          `json:"UpdatedAt" bson:"updated_at"`
}
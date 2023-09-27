package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Register struct {
	ID             primitive.ObjectID      `bson:"_id"`
	FirstName     		string             `json:"Firstname" bson:"FirstName" binding:"required"`
	LastName     	 	string             `json:"Lastname" bson:"LastName" binding:"required"`
	Age			  		int                `json:"Age" bson:"Age" binding:"required"`
	Email    	  		string             `json:"Email" bson:"Email" binding:"required"`
	Password      		string             `json:"Password" bson:"Password" binding:"required,min=8"`
	PasswordConfirm     string    	 	   `json:"PasswordConfirm" bson:"PasswordConfirm,omitempty" binding:"required"`
	CreatedAt           time.Time 		   `json:"CreatedAt" bson:"CreatedAt"`
	UpdatedAt           time.Time           `json:"UpdatedAt" bson:"UpdatedAt"`
}

type Login struct{
	Email              string    `json:"Email" bson:"Email" binding:"required"`
	Password           string    `json:"Password" bson:"Password" binding:"required,min=8"`
}
type LoginResponse struct {
	Response           string    `json:"Response"`
}
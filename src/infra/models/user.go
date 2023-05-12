package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	Street  string `bson:"street,omitempty"`
	City    string `bson:"city,omitempty"`
	State   string `bson:"state,omitempty"`
	Country string `bson:"country,omitempty"`
}

type User struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	BirthDate   string             `bson:"birthDate,omitempty"`
	Email       string             `bson:"email,omitempty"`
	Password    string             `bson:"password,omitempty"`
	Address     *Address           `bson:"address,omitempty"`
	CreatedDate time.Time          `bson:"createdDate,omitempty"`
	UpdatedDate time.Time          `bson:"updatedDate,omitempty"`
}

package models

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStatus string

const (
	StatusActive   UserStatus = "active"
	StatusInactive UserStatus = "inactive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Email       string             `bson:"email" json:"email"`
	PhoneNumber string             `bson:"phone_number" json:"phone_number"`
	BusinessID  primitive.ObjectID `bson:"business_id" json:"business_id"`
	Status      UserStatus         `bson:"status" json:"status"`
	Avatar      string             `bson:"avatar" json:"avatar"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

func (obj User) Validate() error {
	return validation.ValidateStruct(&obj,
		validation.Field(&obj.Name, validation.Required),
		validation.Field(&obj.BusinessID, validation.Required),
	)
}

type UserSearch struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty"`
}

func (obj UserSearch) GetUserSearchBSONObj() bson.M {
	self := bson.M{}
	if obj.ID != primitive.NilObjectID {
		self["_id"] = obj.ID
	}

	if obj.Name != "" {
		regexPattern := fmt.Sprintf(".*%s.*", obj.Name)
		self["name"] = bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: regexPattern, Options: "i"}}}
	}

	return self
}

package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BusinessStatus string

const (
	BusinessStatusActive    BusinessStatus = "active"
	BusinessStatusSuspended BusinessStatus = "suspended"
	BusinessStatusTrial     BusinessStatus = "trial"
)

type Business struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Email       string             `bson:"email" json:"email"`
	PhoneNumber string             `bson:"phone_number" json:"phone_number"`
	Status      BusinessStatus     `bson:"status" json:"status"` // active, suspended, trial
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

func (obj Business) Validate() error {
	return validation.ValidateStruct(&obj,
		validation.Field(&obj.Name, validation.Required),
		validation.Field(&obj.Email, validation.Required),
		validation.Field(&obj.PhoneNumber, validation.Required),
	)
}

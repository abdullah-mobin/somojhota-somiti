package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BalanceSheet struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BusinessID primitive.ObjectID `bson:"business_id" json:"business_id"`
	Total      float64            `bson:"total" json:"total"`
	Month      int                `bson:"month" json:"month"`
	Year       int                `bson:"year" json:"year"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

func (obj BalanceSheet) Validate() error {
	return validation.ValidateStruct(&obj,
		validation.Field(&obj.BusinessID, validation.Required),
	)
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BusinessID primitive.ObjectID `bson:"business_id" json:"business_id"`
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id"`
	Amount     float64            `bson:"amount" json:"amount"`
	Balance    float64            `bson:"balance" json:"balance"`
	Date       string             `bson:"date" json:"date"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

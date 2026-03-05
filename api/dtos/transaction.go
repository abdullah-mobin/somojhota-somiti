package dtos

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateTransactionDTO struct {
	BusinessID string  `json:"business_id"`
	UserID     string  `json:"user_id"`
	Date       string  `json:"date"`
	Amount     float64 `json:"amount"`
	Balance    float64 `json:"balance"`
}

type UpdateTransactionDTO struct {
	Date    *string  `json:"date"`
	Amount  *float64 `json:"amount"`
	Balance *float64 `json:"balance"`
}

func (obj CreateTransactionDTO) Validate() error {
	return validation.ValidateStruct(&obj,
		validation.Field(&obj.BusinessID, validation.Required),
		validation.Field(&obj.UserID, validation.Required),
		validation.Field(&obj.Date, validation.Required),
		validation.Field(&obj.Amount, validation.Required),
		validation.Field(&obj.Balance, validation.Required),
	)
}

func (obj UpdateTransactionDTO) Validate() error {
	if obj.Amount == nil && obj.Balance == nil && obj.Date == nil {
		return errors.New("at least one field is required")
	}
	if obj.Amount != nil && *obj.Amount < 0 {
		return errors.New("amount cannot be negative")
	}
	if obj.Balance != nil && *obj.Balance < 0 {
		return errors.New("balance cannot be negative")
	}
	return nil
}

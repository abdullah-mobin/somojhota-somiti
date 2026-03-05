package dtos

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateCostDTO struct {
	BusinessID  string  `json:"business_id"`
	UserID      string  `json:"user_id"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Balance     float64 `json:"balance"`
}

type UpdateCostDTO struct {
	Description *string  `json:"description"`
	Date        *string  `json:"date"`
	Balance     *float64 `json:"balance"`
}

func (obj CreateCostDTO) Validate() error {
	if obj.Balance < 0 {
		return errors.New("invalid balance: cannot be negative")
	}
	return validation.ValidateStruct(&obj,
		validation.Field(&obj.BusinessID, validation.Required),
		validation.Field(&obj.UserID, validation.Required),
		validation.Field(&obj.Date, validation.Required),
		validation.Field(&obj.Balance, validation.Required),
	)
}

func (obj UpdateCostDTO) Validate() error {
	if obj.Description == nil && obj.Balance == nil && obj.Date == nil {
		return errors.New("at least one field is required")
	}
	if obj.Balance != nil && *obj.Balance < 0 {
		return errors.New("balance cannot be negative")
	}
	return nil
}

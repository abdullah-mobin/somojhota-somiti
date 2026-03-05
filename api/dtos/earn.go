package dtos

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateEarnDTO struct {
	BusinessID  string  `json:"business_id"`
	UserID      string  `json:"user_id"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Balance     float64 `json:"balance"`
}

type UpdateEarnDTO struct {
	Description *string  `json:"description"`
	Date        *string  `json:"date"`
	Balance     *float64 `json:"balance"`
}

func (obj CreateEarnDTO) Validate() error {
	return validation.ValidateStruct(&obj,
		validation.Field(&obj.BusinessID, validation.Required),
		validation.Field(&obj.UserID, validation.Required),
		validation.Field(&obj.Date, validation.Required),
		validation.Field(&obj.Balance, validation.Required),
	)
}

func (obj UpdateEarnDTO) Validate() error {
	if obj.Description == nil && obj.Balance == nil && obj.Date == nil {
		return errors.New("at least one field is required")
	}
	if obj.Balance != nil && *obj.Balance < 0 {
		return errors.New("balance cannot be negative")
	}
	return nil
}

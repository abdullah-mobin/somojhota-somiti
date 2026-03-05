package dtos

import validation "github.com/go-ozzo/ozzo-validation"

type LoginDTO struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (obj LoginDTO) Validate() error {
	return validation.ValidateStruct(&obj,
		validation.Field(&obj.PhoneNumber, validation.Required),
		validation.Field(&obj.Password, validation.Required),
	)
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refresh_token"`
}

func (obj RefreshTokenDTO) Validate() error {
	return validation.ValidateStruct(&obj,
		validation.Field(&obj.RefreshToken, validation.Required),
	)
}

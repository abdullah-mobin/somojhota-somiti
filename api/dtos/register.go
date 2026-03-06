package dtos

import validation "github.com/go-ozzo/ozzo-validation"

type RegisterBusinessDTO struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterUserDTO struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	BusinessID  string `json:"business_id"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (obj RegisterBusinessDTO) Validate() error {
	return validation.ValidateStruct(&obj,
		validation.Field(&obj.Name, validation.Required),
		validation.Field(&obj.PhoneNumber, validation.Required),
	)
}
func (obj RegisterUserDTO) Validate() error {
	return validation.ValidateStruct(&obj,
		validation.Field(&obj.Name, validation.Required),
		validation.Field(&obj.PhoneNumber, validation.Required),
		validation.Field(&obj.BusinessID, validation.Required),
		validation.Field(&obj.Password, validation.Required),
	)
}

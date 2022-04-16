package apimodelCustomer

import "gopkg.in/go-playground/validator.v9"

type CustomerProfileUpdateRequest struct {
	CustomerId  string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	LineId      string `json:"line_id" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type CustomerProfileUpdateResponse struct {
}

func (o *CustomerProfileUpdateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

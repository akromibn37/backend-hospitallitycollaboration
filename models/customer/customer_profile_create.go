package apimodelCustomer

import "gopkg.in/go-playground/validator.v9"

type CustomerProfileCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	LineId      string `json:"line_id" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	ServiceId   string `json:"svc_id" validate:"required"`
}

type CustomerProfileCreateResponse struct {
}

func (o *CustomerProfileCreateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

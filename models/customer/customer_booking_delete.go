package apimodelCustomer

import "gopkg.in/go-playground/validator.v9"

type CustomerBookingDeleteRequest struct {
	CustomerBookingId string `json:"id" validate:"required"`
}

type CustomerBookingDeleteResponse struct {
}

func (o *CustomerBookingDeleteRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

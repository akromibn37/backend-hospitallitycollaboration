package apimodelCustomer

import "gopkg.in/go-playground/validator.v9"

type CustomerBookingGetRequest struct {
	CustomerServiceId string `json:"id" validate:"required"`
}

type CustomerBookingGetResponse struct {
	Detail []CustomerBookingDetail `json:"Detail"`
}

func (o *CustomerBookingGetRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

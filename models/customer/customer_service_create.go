package apimodelCustomer

import "gopkg.in/go-playground/validator.v9"

type CustomerServiceCreateRequest struct {
	CustomerId string  `json:"customer_id" validate:"required"`
	ServiceId  *string `json:"svc_id"`
	UserId     *string `json:"user_id"`
	Status     *string `json:"status"`
	StartDate  *string `json:"start_date"`
	EndDate    *string `json:"end_date"`
}

type CustomerServiceCreateResponse struct {
}

func (o *CustomerServiceCreateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

package apimodelCustomer

import "gopkg.in/go-playground/validator.v9"

type CustomerServiceUpdateRequest struct {
	CustomerServiceId string  `json:"id" validate:"required"`
	CustomerId        string  `json:"customer_id"`
	ServiceId         *string `json:"svc_id"`
	HospitalityId     *string `json:"hos_id"`
	UserId            *string `json:"user_id"`
	Status            *string `json:"status"`
}

type CustomerServiceUpdateResponse struct {
}

func (o *CustomerServiceUpdateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

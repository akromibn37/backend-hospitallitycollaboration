package apimodelCustomer

import "gopkg.in/go-playground/validator.v9"

type CustomerServiceCreateRequest struct {
	CustomerId    string  `json:"customer_id"`
	ServiceId     *string `json:"svc_id"`
	HospitalityId *string `json:"hos_id"`
	UserId        *string `json:"user_id"`
	Status        *string `json:"status"`
}

type CustomerServiceCreateResponse struct {
}

func (o *CustomerServiceCreateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

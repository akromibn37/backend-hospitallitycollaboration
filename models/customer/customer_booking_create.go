package apimodelCustomer

import "gopkg.in/go-playground/validator.v9"

type CustomerBookingCreateRequest struct {
	CustomerServiceId string `json:"cus_svc_id" validate:"required"`
	HospitalityId     string `json:"hos_id"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
}

type CustomerBookingCreateResponse struct {
}

func (o *CustomerBookingCreateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

package apimodelCustomer

import "gopkg.in/go-playground/validator.v9"

type CustomerBookingUpdateRequest struct {
	CustomerBookingId string `json:"id" validate:"required"`
	HospitalityId     string `json:"hos_id"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
}

type CustomerBookingUpdateResponse struct {
}

func (o *CustomerBookingUpdateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

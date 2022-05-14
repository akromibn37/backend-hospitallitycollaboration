package apimodelCustomer

import "gopkg.in/go-playground/validator.v9"

type CustomerBookingGetAllRequest struct {
}

type CustomerBookingGetAllResponse struct {
	Detail []CustomerBookingDetail `json:"Detail"`
}

type CustomerBookingDetail struct {
	CustomerBookingId        string `json:"id"`
	CustomerServiceId        string `json:"cus_svc_id"`
	HospitalityId            string `json:"hos_id"`
	StartDate                string `json:"start_date"`
	EndDate                  string `json:"end_date"`
	HospitalityName          string `json:"hos_name"`
	HospitalityContactName   string `json:"hos_contact_name"`
	HospitalityContactNumber string `json:"hos_phone_number"`
}

func (o *CustomerBookingGetAllRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

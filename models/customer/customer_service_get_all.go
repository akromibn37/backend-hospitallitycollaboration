package apimodelCustomer

import "gopkg.in/go-playground/validator.v9"

type CustomerServiceGetAllRequest struct {
}

type CustomerServiceGetAllResponse struct {
	Detail []CustomerServiceDetail `json:"Detail"`
}

type CustomerServiceDetail struct {
	CustomerServiceId        string  `json:"id"`
	CustomerName             *string `json:"customer_name"`
	CustomerPhoneNumber      *string `json:"customer_phone_number"`
	HospitalityName          *string `json:"hos_name"`
	HospitalityContactName   *string `json:"hos_contact_name"`
	HospitalityContactNumber *string `json:"hos_phone_number"`
	ServiceName              *string `json:"svc_name"`
	UserName                 *string `json:"user_name"`
	Status                   *string `json:"status"`
	CustomerId               string  `json:"customer_id"`
	ServiceId                *string `json:"svc_id"`
	UserId                   *string `json:"user_id"`
	HospitalityId            *string `json:"hos_id"`
}

func (o *CustomerServiceGetAllRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

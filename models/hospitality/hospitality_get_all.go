package apimodelHospitality

import "gopkg.in/go-playground/validator.v9"

type HospitalityGetAllRequest struct {
}

type HospitalityGetAllResponse struct {
	Detail []HospitalityDetail `json:"Detail"`
}

type HospitalityDetail struct {
	HospitalityId          string `json:"id"`
	HospitalityName        string `json:"hos_name"`
	HospitalityDescription string `json:"hos_desc"`
	HospitalityType        string `json:"hos_type"`
	HospitalityContactName string `json:"hos_contact_name"`
	HospitalityPhoneNumber string `json:"hos_phone_number"`
}

func (o *HospitalityGetAllRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

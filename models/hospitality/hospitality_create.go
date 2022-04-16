package apimodelHospitality

import "gopkg.in/go-playground/validator.v9"

type HospitalityCreateRequest struct {
	HospitalityName        string `json:"hos_name" validate:"required"`
	HospitalityDescription string `json:"hos_desc" validate:"required"`
	HospitalityType        string `json:"hos_type" validate:"required"`
	HospitalityContactName string `json:"hos_contact_name" validate:"required"`
	HospitalityPhoneNumber string `json:"hos_phone_number" validate:"required"`
}

type HospitalityCreateResponse struct {
}

func (o *HospitalityCreateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

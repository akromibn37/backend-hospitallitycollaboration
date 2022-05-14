package apimodelHospitality

import "gopkg.in/go-playground/validator.v9"

type HospitalityTypeCreateRequest struct {
	HospitalityTypeName string `json:"hos_type_name" validate:"required"`
}

type HospitalityTypeCreateResponse struct {
}

func (o *HospitalityTypeCreateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

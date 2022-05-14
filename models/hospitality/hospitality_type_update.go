package apimodelHospitality

import "gopkg.in/go-playground/validator.v9"

type HospitalityTypeUpdateRequest struct {
	HospitalityTypeId   string `json:"id" validate:"required"`
	HospitalityTypeName string `json:"hos_type_name" validate:"required"`
}

type HospitalityTypeUpdateResponse struct {
}

func (o *HospitalityTypeUpdateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

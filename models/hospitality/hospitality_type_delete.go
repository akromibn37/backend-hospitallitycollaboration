package apimodelHospitality

import "gopkg.in/go-playground/validator.v9"

type HospitalityTypeDeleteRequest struct {
	HospitalityTypeId string `json:"id" validate:"required"`
}

type HospitalityTypeDeleteResponse struct {
}

func (o *HospitalityTypeDeleteRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

package apimodelHospitality

import "gopkg.in/go-playground/validator.v9"

type HospitalityTypeGetAllRequest struct {
}

type HospitalityTypeGetAllResponse struct {
	Detail []HospitalityTypeDetail `json:"Detail"`
}

type HospitalityTypeDetail struct {
	HospitalityTypeId   string `json:"id"`
	HospitalityTypeName string `json:"hos_type_name"`
}

func (o *HospitalityTypeGetAllRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

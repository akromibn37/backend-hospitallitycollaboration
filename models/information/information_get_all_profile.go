package apimodelInformation

import "gopkg.in/go-playground/validator.v9"

type InformationGetAllProfileRequest struct {
	// UserId string `json:"userId" validate:"required,max=45"`
}

type InformationGetAllProfileResponse struct {
	Profile []Profile `json:"data"`
}

func (o *InformationGetAllProfileRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

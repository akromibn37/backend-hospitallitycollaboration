package apimodelService

import "gopkg.in/go-playground/validator.v9"

type ServiceDeleteRequest struct {
	ServiceId string `json:"id" validate:"required,max=45"`
}

type ServiceDeleteResponse struct {
}

func (o *ServiceDeleteRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

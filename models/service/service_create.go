package apimodelService

import "gopkg.in/go-playground/validator.v9"

type ServiceCreateRequest struct {
	ServiceName        string `json:"svc_name" validate:"required,max=45"`
	ServiceDescription string `json:"svc_desc"`
}

type ServiceCreateResponse struct {
}

func (o *ServiceCreateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

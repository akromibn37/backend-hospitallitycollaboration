package apimodelService

import "gopkg.in/go-playground/validator.v9"

type ServiceUpdateRequest struct {
	ServiceId          string `json:"id" validate:"required,max=45"`
	ServiceName        string `json:"svc_name" validate:"required,max=45"`
	ServiceDescription string `json:"svc_desc" validate:"required"`
}

type ServiceUpdateResponse struct {
}

func (o *ServiceUpdateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

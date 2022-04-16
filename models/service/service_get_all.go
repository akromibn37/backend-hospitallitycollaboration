package apimodelService

import "gopkg.in/go-playground/validator.v9"

type ServiceGetAllRequest struct {
}

type ServiceGetAllResponse struct {
	Detail []ServiceDetail `json:"Detail"`
}

type ServiceDetail struct {
	ServiceId          string `json:"id"`
	ServiceName        string `json:"svc_name"`
	ServiceDescription string `json:"svc_desc"`
}

func (o *ServiceGetAllRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

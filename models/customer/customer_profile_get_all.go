package apimodelCustomer

import "gopkg.in/go-playground/validator.v9"

type CustomerProfileGetAllRequest struct {
}

type CustomerProfileGetAllResponse struct {
	Detail []CustomerProfileDetail `json:"Detail"`
}

type CustomerProfileDetail struct {
	CustomerId  string `json:"id"`
	Name        string `json:"name"`
	LineId      string `json:"line_id"`
	PhoneNumber string `json:"phone_number"`
	Nationality string `json:"nationality"`
}

func (o *CustomerProfileGetAllRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

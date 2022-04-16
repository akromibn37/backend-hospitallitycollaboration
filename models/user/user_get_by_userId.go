package apimodelUser

import "gopkg.in/go-playground/validator.v9"

type UserGetByUserIdRequest struct {
	UserId string `json:"userId" validate:"required,max=45"`
}

type UserGetByUserIdResponse struct {
	UserId   string   `json:"userId"`
	UserType string   `json:"role"`
	UserData UserData `json:"data"`
}

func (o *UserGetByUserIdRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

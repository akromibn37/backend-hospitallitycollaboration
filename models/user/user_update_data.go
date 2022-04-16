package apimodelUser

import "gopkg.in/go-playground/validator.v9"

type UserUpdateDataRequest struct {
	UserId      string  `json:"id" validate:"required,max=45"`
	UserType    *string `json:"userType"`
	DisplayName *string `json:"displayName"`
	Active      *bool   `json:"active"`
}

type UserUpdateDataResponse struct {
}

func (o *UserUpdateDataRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

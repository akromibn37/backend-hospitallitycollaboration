package apimodelUser

import "gopkg.in/go-playground/validator.v9"

type UserCreateRequest struct {
	Username    string `json:"username" validate:"required,max=45"`
	PassWord    string `json:"password" validate:"required"`
	UserType    string `json:"userType" validate:"required,max=10"`
	DisplayName string `json:"displayName" validate:"required,max=45"`
	CreateBy    string `json:"createBy" validate:"required,max=50"`
	UpdateBy    string `json:"updateBy" validate:"required,max=50"`
}

type UserCreateResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (o *UserCreateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

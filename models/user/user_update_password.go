package apimodelUser

import "gopkg.in/go-playground/validator.v9"

type UserUpdatePasswordRequest struct {
	UserId      string `json:"userId" validate:"required,max=45"`
	PassWord    string `json:"password" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

type UserUpdatePasswordResponse struct {
}

func (o *UserUpdatePasswordRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

package apimodelUser

import "gopkg.in/go-playground/validator.v9"

type UserLoginRequest struct {
	Username string `json:"username" validate:"required,max=45"`
	PassWord string `json:"password" validate:"required"`
}

type UserLoginResponse struct {
	UserId   string   `json:"userId"`
	UserType string   `json:"role"`
	UserData UserData `json:"data"`
}

type UserData struct {
	DisplayName string  `json:"displayName"`
	PhotoUrl    string  `json:"photoURL"`
	Setting     Setting `json:"settings"`
	Shortcut    []string
}

type Setting struct {
}

func (o *UserLoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

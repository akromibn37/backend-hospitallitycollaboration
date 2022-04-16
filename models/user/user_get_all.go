package apimodelUser

import "gopkg.in/go-playground/validator.v9"

type UserGetAllRequest struct {
}

type UserGetAllResponse struct {
	Data []UserAuth `json:"data"`
}

type UserAuth struct {
	UserId      string `json:"id"`
	Username    string `json:"username"`
	UserType    string `json:"userType"`
	DisplayName string `json:"displayName"`
	Active      string `json:"active"`
	CreateDate  string `json:"createDate"`
	UpdateDate  string `json:"updateDate"`
}

func (o *UserGetAllRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

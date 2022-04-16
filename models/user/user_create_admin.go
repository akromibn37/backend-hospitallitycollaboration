package apimodelUser

import "gopkg.in/go-playground/validator.v9"

type UserCreateByAdminRequest struct {
	Username    string  `json:"username" validate:"required,max=45"`
	PassWord    string  `json:"password" validate:"required"`
	UserType    string  `json:"usertype" validate:"required,max=10"`
	DisplayName string  `json:"displayName" validate:"required,max=45"`
	TitleName   *string `json:"title_name"`
	Name        string  `json:"name"`
	Surname     string  `json:"surname"`
	DateOfBirth string  `json:"date_of_birth"`
	Identifier  string  `json:"identifier"`
	PhoneNumber string  `json:"phone_number"`
	Email       *string `json:"email"`
	Addr        string  `json:"addr"`
	SubDistrict string  `json:"sub_district"`
	District    string  `json:"district"`
	Province    string  `json:"province"`
	PostalCode  string  `json:"postal_code"`
	CreateBy    string  `json:"createBy"`
	UpdateBy    string  `json:"updateBy"`
}

type UserCreateByAdminResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (o *UserCreateByAdminRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

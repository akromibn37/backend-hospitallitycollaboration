package apimodelInformation

import "gopkg.in/go-playground/validator.v9"

type InformationGetProfileRequest struct {
	UserId string `json:"id" validate:"required,max=45"`
}

type InformationGetProfileResponse struct {
	Profile []Profile `json:"data"`
}
type Profile struct {
	UserId      string  `json:"id" validate:"required,max=45"`
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
}

func (o *InformationGetProfileRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

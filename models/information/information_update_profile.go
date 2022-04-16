package apimodelInformation

import "gopkg.in/go-playground/validator.v9"

type InformationUpdateProfileRequest struct {
	UserId      string  `json:"id" validate:"required,max=45"`
	Name        *string `json:"name"`
	Surname     *string `json:"surname"`
	DateOfBirth *string `json:"date_of_birth"`
	Identifier  *string `json:"identifier"`
	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`
	Addr        *string `json:"addr"`
	SubDistrict *string `json:"sub_district"`
	District    *string `json:"district"`
	Province    *string `json:"province"`
	PostalCode  *string `json:"postal_code"`
}

type InformationUpdateProfileResponse struct {
}

func (o *InformationUpdateProfileRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

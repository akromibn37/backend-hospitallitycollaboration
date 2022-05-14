package service

import (
	"time"

	databaseUser "github.com/akromibn37/hospitalityCollaboration/dao/database/users"
	apimodelInformation "github.com/akromibn37/hospitalityCollaboration/models/information"
	"github.com/akromibn37/hospitalityCollaboration/pkg/logging"
)

type Information struct {
}

func (i *Information) InformationGetProfile(req *apimodelInformation.InformationGetProfileRequest, t *logging.Timelog) (res *apimodelInformation.InformationGetProfileResponse, err error) {

	t.TimeInDb = time.Now()
	data, err := databaseUser.GetProfile(req.UserId)
	if err != nil {
		return nil, err
	}
	t.TimeOutDb = time.Now()
	resData := make([]apimodelInformation.Profile, 1)
	resData[0].UserId = data.UserId
	resData[0].Name = data.Name
	resData[0].Surname = data.Surname
	resData[0].DateOfBirth = data.DateOfBirth
	resData[0].Identifier = data.Identifier
	resData[0].PhoneNumber = data.PhoneNumber
	resData[0].Email = data.Email
	resData[0].Addr = data.Addr
	resData[0].SubDistrict = data.SubDistrict
	resData[0].District = data.District
	resData[0].Province = data.Province
	resData[0].PostalCode = data.PostalCode
	res = &apimodelInformation.InformationGetProfileResponse{Profile: resData}

	return res, nil
}

func (i *Information) InformationGetAllProfile(t *logging.Timelog) (res *apimodelInformation.InformationGetAllProfileResponse, err error) {

	t.TimeInDb = time.Now()
	data, err := databaseUser.GetAllProfile()
	if err != nil {
		return nil, err
	}
	t.TimeOutDb = time.Now()
	resData := make([]apimodelInformation.Profile, len(data))
	for i := 0; i < len(data); i++ {
		resData[i].UserId = data[i].UserId
		resData[i].Name = data[i].Name
		resData[i].Surname = data[i].Surname
		resData[i].DateOfBirth = data[i].DateOfBirth
		resData[i].Identifier = data[i].Identifier
		resData[i].PhoneNumber = data[i].PhoneNumber
		resData[i].Email = data[i].Email
		resData[i].Addr = data[i].Addr
		resData[i].SubDistrict = data[i].SubDistrict
		resData[i].District = data[i].District
		resData[i].Province = data[i].Province
		resData[i].PostalCode = data[i].PostalCode
	}
	res = &apimodelInformation.InformationGetAllProfileResponse{Profile: resData}

	return res, nil
}

func (i *Information) InformationUpdateProfile(req *apimodelInformation.InformationUpdateProfileRequest, t *logging.Timelog) (res *apimodelInformation.InformationUpdateProfileResponse, err error) {
	t.TimeInDb = time.Now()
	data, err := databaseUser.GetProfile(req.UserId)
	if err != nil {
		return nil, err
	}

	if req.Name == nil {
		req.Name = &data.Name
	}
	if req.Surname == nil {
		req.Surname = &data.Surname
	}
	if req.DateOfBirth == nil {
		req.DateOfBirth = &data.DateOfBirth
	}
	if req.Identifier == nil {
		req.Identifier = &data.Identifier
	}
	if req.PhoneNumber == nil {
		req.PhoneNumber = &data.PhoneNumber
	}
	if req.Email == nil {
		req.Email = data.Email
	}
	if req.Addr == nil {
		req.Addr = &data.Addr
	}
	if req.SubDistrict == nil {
		req.SubDistrict = &data.SubDistrict
	}
	if req.District == nil {
		req.District = &data.District
	}
	if req.Province == nil {
		req.Province = &data.Province
	}
	if req.PostalCode == nil {
		req.PostalCode = &data.PostalCode
	}

	err = databaseUser.UpdateProfile(req)

	t.TimeOutDb = time.Now()
	res = &apimodelInformation.InformationUpdateProfileResponse{}

	return res, nil
}

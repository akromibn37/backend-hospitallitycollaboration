package service

import (
	"fmt"
	"time"

	"github.com/akromibn37/hospitalityCollaboration/constant"
	databaseUser "github.com/akromibn37/hospitalityCollaboration/dao/database/users"
	apimodelInformation "github.com/akromibn37/hospitalityCollaboration/models/information"
	apimodelUser "github.com/akromibn37/hospitalityCollaboration/models/user"
	"github.com/akromibn37/hospitalityCollaboration/pkg/logging"
	"github.com/akromibn37/hospitalityCollaboration/pkg/util"
	projutil "github.com/akromibn37/hospitalityCollaboration/util"
)

type User struct {
}

func (o *User) UserCreate(req *apimodelUser.UserCreateRequest, t *logging.Timelog) (res *apimodelUser.UserCreateResponse, err error) {
	userId := projutil.GenerateID(constant.USER_ID)
	req.PassWord = util.EncodeMD5(req.PassWord)
	// userProfile := &apimodelInformation.InformationUpdateProfileRequest{
	// 	UserId:           userId,
	// 	TitleName:        nil,
	// 	Name:             "",
	// 	Surname:          nil,
	// 	DateOfBirth:      nil,
	// 	Identifier:       nil,
	// 	PhoneNumber:      nil,
	// 	Email:            nil,
	// 	Addr:             nil,
	// 	SubDistrict:      nil,
	// 	District:         nil,
	// 	Province:         nil,
	// 	PostalCode:       nil,
	// 	FatherName:       nil,
	// 	MotherName:       nil,
	// 	FatherBirthDate:  nil,
	// 	MotherBirthDate:  nil,
	// 	LivingStatus:     nil,
	// 	FatherOccupation: nil,
	// 	MotherOccupation: nil,
	// 	HouseHoldIncome:  nil,
	// 	IsChildren:       nil,
	// 	SchoolName:       nil,
	// }
	t.TimeInDb = time.Now()
	err = databaseUser.CreateBasicUser(req, userId)
	if err != nil {
		return nil, err
	}
	err = databaseUser.CreateProfileByRegister(userId, req)
	if err != nil {
		return nil, err
	}
	t.TimeOutDb = time.Now()
	fmt.Println("err:", err)

	res = &apimodelUser.UserCreateResponse{
		Code:        "0000",
		Description: "Success",
	}

	return res, nil
}

func (o *User) UserCreateByAdmin(req *apimodelUser.UserCreateByAdminRequest, t *logging.Timelog) (res *apimodelUser.UserCreateByAdminResponse, err error) {
	userId := projutil.GenerateID(constant.USER_ID)
	req.PassWord = util.EncodeMD5(req.PassWord)
	t.TimeInDb = time.Now()
	userCreateRequest := &apimodelUser.UserCreateRequest{
		Username:    req.Username,
		PassWord:    req.PassWord,
		UserType:    req.UserType,
		DisplayName: req.DisplayName,
		CreateBy:    req.CreateBy,
		UpdateBy:    req.UpdateBy,
	}
	err = databaseUser.CreateBasicUser(userCreateRequest, userId)

	if err != nil {
		return nil, err
	}

	userCreateProfile := &apimodelInformation.InformationUpdateProfileRequest{
		UserId:      userId,
		Name:        &req.Name,
		Surname:     &req.Surname,
		DateOfBirth: &req.DateOfBirth,
		Identifier:  &req.Identifier,
		PhoneNumber: &req.PhoneNumber,
		Email:       req.Email,
		Addr:        &req.Addr,
		SubDistrict: &req.SubDistrict,
		District:    &req.District,
		Province:    &req.Province,
		PostalCode:  &req.PostalCode,
	}

	err = databaseUser.CreateProfile(userCreateProfile)
	if err != nil {
		return nil, err
	}

	t.TimeOutDb = time.Now()
	res = &apimodelUser.UserCreateByAdminResponse{
		Code:        "0000",
		Description: "Success",
	}

	return res, nil
}

func (o *User) UserLogin(req *apimodelUser.UserLoginRequest, t *logging.Timelog) (res *apimodelUser.UserLoginResponse, err error) {
	t.TimeInDb = time.Now()
	req.PassWord = util.EncodeMD5(req.PassWord)
	data, err := databaseUser.Login(req)
	if err != nil {
		return nil, err
	}
	t.TimeOutDb = time.Now()
	res = &apimodelUser.UserLoginResponse{}
	res.UserId = data.UserId
	res.UserData.DisplayName = data.DisplayName
	res.UserType = data.UserType

	return res, nil
}

func (o *User) UserGetByUserId(req *apimodelUser.UserGetByUserIdRequest, t *logging.Timelog) (res *apimodelUser.UserGetByUserIdResponse, err error) {
	t.TimeInDb = time.Now()
	data, err := databaseUser.GetByUserId(req)
	if err != nil {
		return nil, err
	}
	t.TimeOutDb = time.Now()
	res = &apimodelUser.UserGetByUserIdResponse{}
	res.UserId = data.UserId
	res.UserData.DisplayName = data.DisplayName
	res.UserType = data.UserType

	return res, nil
}

func (o *User) UserUpdatePassword(req *apimodelUser.UserUpdatePasswordRequest, t *logging.Timelog) (res *apimodelUser.UserUpdatePasswordResponse, err error) {
	t.TimeInDb = time.Now()
	req.PassWord = util.EncodeMD5(req.PassWord)
	req.NewPassword = util.EncodeMD5(req.NewPassword)
	existPassword, err := databaseUser.GetPassword(req.UserId)
	if err != nil {
		return nil, err
	}
	if existPassword != req.PassWord {
		return nil, fmt.Errorf("UserId or Password Not match")
	}
	err = databaseUser.UpdatePassword(req)
	if err != nil {
		return nil, err
	}
	t.TimeOutDb = time.Now()
	res = &apimodelUser.UserUpdatePasswordResponse{}

	return res, nil
}

func (o *User) UserUpdateData(req *apimodelUser.UserUpdateDataRequest, t *logging.Timelog) (res *apimodelUser.UserUpdateDataResponse, err error) {
	t.TimeInDb = time.Now()
	data, err := databaseUser.GetData(req.UserId)
	if err != nil {
		return nil, err
	}
	if req.UserType == nil {
		req.UserType = &data.UserType
	}
	if req.DisplayName == nil {
		req.DisplayName = &data.DisplayName
	}
	if req.Active == nil {
		req.Active = &data.Active
	}

	err = databaseUser.UpdateData(req)
	if err != nil {
		return nil, err
	}

	t.TimeOutDb = time.Now()
	res = &apimodelUser.UserUpdateDataResponse{}

	return res, nil
}

func (o *User) UserGetAll(t *logging.Timelog) (res *apimodelUser.UserGetAllResponse, err error) {

	t.TimeInDb = time.Now()
	data, err := databaseUser.GetAllData()
	if err != nil {
		return nil, err
	}
	t.TimeOutDb = time.Now()
	resData := make([]apimodelUser.UserAuth, len(data))
	for i := 0; i < len(data); i++ {
		active := "active"
		if data[i].Active == true {
			active = "active"
		} else {
			active = "inactive"
		}
		resData[i].UserId = data[i].UserId
		resData[i].Username = data[i].UserName
		resData[i].UserType = data[i].UserType
		resData[i].DisplayName = data[i].DisplayName
		resData[i].Active = active
		resData[i].CreateDate = data[i].CreateDate
		resData[i].UpdateDate = data[i].UpdateDate
	}
	res = &apimodelUser.UserGetAllResponse{Data: resData}

	return res, nil
}

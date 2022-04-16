package databaseUser

import (
	"context"
	"fmt"
	"log"

	database "github.com/akromibn37/hospitalityCollaboration/dao/database"
	apimodelInformation "github.com/akromibn37/hospitalityCollaboration/models/information"
	"github.com/akromibn37/hospitalityCollaboration/pkg/util"
)

type TUserProfile struct {
	UserId      string  `json:"user_id"`
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
	CreateBy    string  `json:"create_by"`
	CreateDate  string  `json:"create_date"`
	UpdateBy    string  `json:"update_by"`
	UpdateDate  string  `json:"update_date"`
}

//GetProfile
func GetProfile(userId string) (res *TUserProfile, err error) {
	db := database.GetDB()
	statement := "select * from t_user_profile where user_id=?"
	resp := TUserProfile{}

	err = db.QueryRow(statement, userId).Scan(&resp.UserId, &resp.Name, &resp.Surname, &resp.DateOfBirth, &resp.Identifier, &resp.PhoneNumber, &resp.Email, &resp.Addr, &resp.SubDistrict, &resp.District, &resp.Province, &resp.PostalCode, &resp.CreateBy, &resp.CreateDate, &resp.UpdateBy, &resp.UpdateDate)
	if err != nil {
		log.Print("err:", err)
		return nil, err
	}
	return &resp, nil
}

//GetAllProfile
func GetAllProfile() (res []*TUserProfile, err error) {
	db := database.GetDB()
	statement := "select * from t_user_profile"
	resp := make([]*TUserProfile, 0)

	results, err := db.Query(statement)
	if err != nil {
		log.Print("err:", err)
		return nil, err
	}
	for results.Next() {
		tmpResp := new(TUserProfile)
		err = results.Scan(&tmpResp.UserId, &tmpResp.Name, &tmpResp.Surname, &tmpResp.DateOfBirth, &tmpResp.Identifier, &tmpResp.PhoneNumber, &tmpResp.Email, &tmpResp.Addr, &tmpResp.SubDistrict, &tmpResp.District, &tmpResp.Province, &tmpResp.PostalCode, &tmpResp.CreateBy, &tmpResp.CreateDate, &tmpResp.UpdateBy, &tmpResp.UpdateDate)
		if err != nil {
			log.Print("err:", err)
			return nil, err
		}
		resp = append(resp, tmpResp)
	}
	return resp, nil
}

//UpdateProfile
func UpdateProfile(req *apimodelInformation.InformationUpdateProfileRequest) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "UPDATE hospitality.t_user_profile SET NAME=?, SURNAME=?, DATE_OF_BIRTH=?, IDENTIFIER=?, PHONE_NUMBER=?, EMAIL=?, ADDR=?, SUB_DICTRICT=?, DISTRICT=?, PROVINCE=?, POSTAL_CODE=?, UPDATE_DATE=? WHERE user_id=?;"
	_, err = tx.Query(statement, req.Name, req.Surname, req.DateOfBirth, req.Identifier, req.PhoneNumber, req.Email, req.Addr, req.SubDistrict, req.District, req.Province, req.PostalCode, util.TimeNow(), req.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

//CreateProfile
func CreateProfile(req *apimodelInformation.InformationUpdateProfileRequest) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	statement := "INSERT INTO hospitality.t_user_profile (user_id, NAME, SURNAME, DATE_OF_BIRTH, IDENTIFIER, PHONE_NUMBER, EMAIL, ADDR, SUB_DICTRICT, DISTRICT, PROVINCE, POSTAL_CODE, CREATE_BY, CREATE_DATE, UPDATE_BY, UPDATE_DATE) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	_, err = tx.Query(statement, req.UserId, req.Name, req.Surname, req.DateOfBirth, req.Identifier, req.PhoneNumber, req.Email, req.Addr, req.SubDistrict, req.District, req.Province, req.PostalCode, "SYSTEM", util.TimeNow(), "SYSTEM", util.TimeNow())
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

//CreateProfile
func CreateProfileByRegister(userId string) (err error) {
	fmt.Println("userId:", userId)
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	statement := "INSERT INTO hospitality.t_user_profile (user_id, NAME, SURNAME, DATE_OF_BIRTH, IDENTIFIER, PHONE_NUMBER, EMAIL, ADDR, SUB_DICTRICT, DISTRICT, PROVINCE, POSTAL_CODE,  CREATE_BY, CREATE_DATE, UPDATE_BY, UPDATE_DATE) VALUES(?, '', '', '', '', '', '', '', '', '', '','','', ?, '', ?);"
	_, err = tx.Query(statement, userId, util.TimeNow(), util.TimeNow())
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

package databaseHospitality

import (
	"context"
	"errors"
	"log"

	database "github.com/akromibn37/hospitalityCollaboration/dao/database"
	apimodelHospitality "github.com/akromibn37/hospitalityCollaboration/models/hospitality"
	"github.com/akromibn37/hospitalityCollaboration/pkg/util"
)

type THospitality struct {
	HospitalityId          string `json:"hos_id"`
	HospitalityName        string `json:"hos_name"`
	HospitalityDescription string `json:"hos_description"`
	HospitalityType        string `json:"hos_type"`
	HospitalityContactName string `json:"hos_contact_name"`
	HospitalityPhoneNumber string `json:"hos_phone_number"`
	CreateBy               string `json:"create_by"`
	CreateDate             string `json:"create_date"`
	UpdateBy               string `json:"update_by"`
	UpdateDate             string `json:"update_date"`
}

//GetScholarshipAll
func GetHospitalityAll() (res []*THospitality, err error) {

	db := database.GetDB()
	statement := "select * from t_hospitality"
	resp := make([]*THospitality, 0)

	results, err := db.Query(statement)
	if err != nil {
		log.Print("err:", err)
		return nil, err
	}
	for results.Next() {
		tmpResp := new(THospitality)
		err = results.Scan(&tmpResp.HospitalityId, &tmpResp.HospitalityName, &tmpResp.HospitalityDescription, &tmpResp.HospitalityType, &tmpResp.HospitalityContactName, &tmpResp.HospitalityPhoneNumber, &tmpResp.CreateBy, &tmpResp.CreateDate, &tmpResp.UpdateBy, &tmpResp.UpdateDate)
		if err != nil {
			log.Print("err:", err)
			return nil, err
		}
		resp = append(resp, tmpResp)
	}

	return resp, nil
}

//Add New User
func CreateHospitality(req *apimodelHospitality.HospitalityCreateRequest, hosid string) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "select count(*) from t_hospitality where hos_name=?"
	result := db.QueryRow(statement, req.HospitalityName)
	var count int
	err = result.Scan(&count)
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("Hospitalityname repeated")
	}
	statement = "INSERT INTO hospitality.t_hospitality (HOS_ID, HOS_NAME, HOS_DESCRIPTION, HOS_TYPE, HOS_CONTACT_NAME, HOS_PHONE_NUMBER, CREATE_BY, CREATE_DATE, UPDATE_BY, UPDATE_DATE) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	_, err = tx.Query(statement, hosid, req.HospitalityName, req.HospitalityDescription, req.HospitalityType, req.HospitalityContactName, req.HospitalityPhoneNumber, "SYSTEM", util.TimeNow(), "SYSTEM", util.TimeNow())
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

//Add New User
func UpdateHospitality(req *apimodelHospitality.HospitalityUpdateRequest) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "UPDATE hospitality.t_hospitality SET HOS_NAME=?, HOS_DESCRIPTION=?, HOS_TYPE=?, HOS_CONTACT_NAME=?, HOS_PHONE_NUMBER=?, UPDATE_DATE=? WHERE HOS_ID=?;"
	_, err = tx.Query(statement, req.HospitalityName, req.HospitalityDescription, req.HospitalityType, req.HospitalityContactName, req.HospitalityPhoneNumber, util.TimeNow(), req.HospitalityId)
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

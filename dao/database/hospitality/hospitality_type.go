package databaseHospitality

import (
	"context"
	"errors"
	"log"

	database "github.com/akromibn37/hospitalityCollaboration/dao/database"
	apimodelHospitality "github.com/akromibn37/hospitalityCollaboration/models/hospitality"
	"github.com/akromibn37/hospitalityCollaboration/pkg/util"
)

type THospitalityType struct {
	HospitalityTypeId   string `json:"hos_type_id"`
	HospitalityTypeName string `json:"hos_type_name"`
	CreateBy            string `json:"create_by"`
	CreateDate          string `json:"create_date"`
	UpdateBy            string `json:"update_by"`
	UpdateDate          string `json:"update_date"`
}

//GetScholarshipAll
func GetHospitalityTypeAll() (res []*THospitalityType, err error) {

	db := database.GetDB()
	statement := "select * from t_hospitality_type"
	resp := make([]*THospitalityType, 0)

	results, err := db.Query(statement)
	if err != nil {
		log.Print("err:", err)
		return nil, err
	}
	for results.Next() {
		tmpResp := new(THospitalityType)
		err = results.Scan(&tmpResp.HospitalityTypeId, &tmpResp.HospitalityTypeName, &tmpResp.CreateBy, &tmpResp.CreateDate, &tmpResp.UpdateBy, &tmpResp.UpdateDate)
		if err != nil {
			log.Print("err:", err)
			return nil, err
		}
		resp = append(resp, tmpResp)
	}

	return resp, nil
}

//Add New User
func CreateHospitalityType(req *apimodelHospitality.HospitalityTypeCreateRequest, hostypeid string) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "select count(*) from t_hospitality_type where hos_type_name=?"
	result := db.QueryRow(statement, req.HospitalityTypeName)
	var count int
	err = result.Scan(&count)
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("HospitalityTypeName repeated")
	}
	statement = "INSERT INTO hospitality.t_hospitality_type (HOS_TYPE_ID, HOS_TYPE_NAME, CREATE_BY, CREATE_DATE, UPDATE_BY, UPDATE_DATE) VALUES(?, ?, ?, ?, ?, ?);"
	_, err = tx.Query(statement, hostypeid, req.HospitalityTypeName, "SYSTEM", util.TimeNow(), "SYSTEM", util.TimeNow())
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
func UpdateHospitalityType(req *apimodelHospitality.HospitalityTypeUpdateRequest) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "UPDATE hospitality.t_hospitality_type SET HOS_TYPE_NAME=?, UPDATE_DATE=? WHERE HOS_TYPE_ID=?;"
	_, err = tx.Query(statement, req.HospitalityTypeName, util.TimeNow(), req.HospitalityTypeId)
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
func DeleteHospitalityType(req *apimodelHospitality.HospitalityTypeDeleteRequest) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "DELETE FROM hospitality.t_hospitality_type WHERE HOS_TYPE_ID=?;"
	_, err = tx.Query(statement, req.HospitalityTypeId)
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

package databaseService

import (
	"context"
	"errors"
	"log"

	database "github.com/akromibn37/hospitalityCollaboration/dao/database"
	apimodelService "github.com/akromibn37/hospitalityCollaboration/models/service"
	"github.com/akromibn37/hospitalityCollaboration/pkg/util"
)

type TService struct {
	ServiceId          string `json:"svc_id"`
	ServiceName        string `json:"svc_name"`
	ServiceDescription string `json:"svc_description"`
	CreateBy           string `json:"create_by"`
	CreateDate         string `json:"create_date"`
	UpdateBy           string `json:"update_by"`
	UpdateDate         string `json:"update_date"`
}

//GetScholarshipAll
func GetServiceAll() (res []*TService, err error) {

	db := database.GetDB()
	statement := "select * from t_service"
	resp := make([]*TService, 0)

	results, err := db.Query(statement)
	if err != nil {
		log.Print("err:", err)
		return nil, err
	}
	for results.Next() {
		tmpResp := new(TService)
		err = results.Scan(&tmpResp.ServiceId, &tmpResp.ServiceName, &tmpResp.ServiceDescription, &tmpResp.CreateBy, &tmpResp.CreateDate, &tmpResp.UpdateBy, &tmpResp.UpdateDate)
		if err != nil {
			log.Print("err:", err)
			return nil, err
		}
		resp = append(resp, tmpResp)
	}

	return resp, nil
}

//Add New User
func CreateService(req *apimodelService.ServiceCreateRequest, svcid string) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "select count(*) from t_service where svc_name=?"
	result := db.QueryRow(statement, req.ServiceName)
	var count int
	err = result.Scan(&count)
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("Servicename repeated")
	}
	statement = "INSERT INTO hospitality.t_service (svc_id, SVC_NAME, SVC_DESCRIPTION, CREATE_BY, CREATE_DATE, UPDATE_BY, UPDATE_DATE) VALUES(?, ?, ?, ?, ?, ?, ?);"
	_, err = tx.Query(statement, svcid, req.ServiceName, req.ServiceDescription, "SYSTEM", util.TimeNow(), "SYSTEM", util.TimeNow())
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
func UpdateService(req *apimodelService.ServiceUpdateRequest) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "UPDATE hospitality.t_service SET SVC_NAME=?, SVC_DESCRIPTION=?, UPDATE_DATE=? WHERE svc_id=?;"
	_, err = tx.Query(statement, req.ServiceName, req.ServiceDescription, util.TimeNow(), req.ServiceId)
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

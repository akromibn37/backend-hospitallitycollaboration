package databaseCustomer

import (
	"context"
	"errors"
	"log"

	database "github.com/akromibn37/hospitalityCollaboration/dao/database"
	apimodelCustomer "github.com/akromibn37/hospitalityCollaboration/models/customer"
	"github.com/akromibn37/hospitalityCollaboration/pkg/util"
)

type TCustomerProfile struct {
	CustomerId  string `json:"customer_id"`
	Name        string `json:"name"`
	LineId      string `json:"line_id"`
	PhoneNumber string `json:"phone_number"`
	CreateBy    string `json:"create_by"`
	CreateDate  string `json:"create_date"`
	UpdateBy    string `json:"update_by"`
	UpdateDate  string `json:"update_date"`
}

//GetScholarshipAll
func GetCustomerProfileAll() (res []*TCustomerProfile, err error) {

	db := database.GetDB()
	statement := "select * from t_customer_profile"
	resp := make([]*TCustomerProfile, 0)

	results, err := db.Query(statement)
	if err != nil {
		log.Print("err:", err)
		return nil, err
	}
	for results.Next() {
		tmpResp := new(TCustomerProfile)
		err = results.Scan(&tmpResp.CustomerId, &tmpResp.Name, &tmpResp.LineId, &tmpResp.PhoneNumber, &tmpResp.CreateBy, &tmpResp.CreateDate, &tmpResp.UpdateBy, &tmpResp.UpdateDate)
		if err != nil {
			log.Print("err:", err)
			return nil, err
		}
		resp = append(resp, tmpResp)
	}

	return resp, nil
}

//Add New User
func CreateCustomerProfile(req *apimodelCustomer.CustomerProfileCreateRequest, customerId string) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "select count(*) from t_customer_profile where name=?"
	result := db.QueryRow(statement, req.Name)
	var count int
	err = result.Scan(&count)
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("CustomerProfile Name repeated")
	}
	statement = "INSERT INTO hospitality.t_customer_profile (customer_id, name, line_id, phone_number, CREATE_BY, CREATE_DATE, UPDATE_BY, UPDATE_DATE) VALUES(?, ?, ?, ?, ?, ?, ?, ?);"
	_, err = tx.Query(statement, customerId, req.Name, req.LineId, req.PhoneNumber, "SYSTEM", util.TimeNow(), "SYSTEM", util.TimeNow())
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
func UpdateCustomerProfile(req *apimodelCustomer.CustomerProfileUpdateRequest) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "UPDATE hospitality.t_customer_profile SET name=?, line_id=?, phone_number=?, UPDATE_DATE=? WHERE customer_id=?;"
	_, err = tx.Query(statement, req.Name, req.LineId, req.PhoneNumber, util.TimeNow(), req.CustomerId)
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

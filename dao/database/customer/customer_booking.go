package databaseCustomer

import (
	"context"
	"log"

	database "github.com/akromibn37/hospitalityCollaboration/dao/database"
	apimodelCustomer "github.com/akromibn37/hospitalityCollaboration/models/customer"
	"github.com/akromibn37/hospitalityCollaboration/pkg/util"
)

type TCustomerBooking struct {
	CustomerBookingId string `json:"cus_bk_id"`
	CustomerServiceId string `json:"cus_svc_id"`
	HospitalityId     string `json:"hos_id"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
	CreateBy          string `json:"create_by"`
	CreateDate        string `json:"create_date"`
	UpdateBy          string `json:"update_by"`
	UpdateDate        string `json:"update_date"`
}

type TCustomerBookingJoin struct {
	CustomerBookingId        string `json:"cus_bk_id"`
	CustomerServiceId        string `json:"cus_svc_id"`
	HospitalityId            string `json:"hos_id"`
	StartDate                string `json:"start_date"`
	EndDate                  string `json:"end_date"`
	HospitalityName          string `json:"hos_name"`
	HospitalityContactName   string `json:"hos_contact_name"`
	HospitalityContactNumber string `json:"hos_phone_number"`
}

//GetScholarshipAll
func GetCustomerBookingAll() (res []*TCustomerBookingJoin, err error) {

	db := database.GetDB()
	statement := "select tbk.cus_bk_id ,tbk.cus_svc_id ,tbk.hos_id ,tbk.start_date ,tbk.end_date ,th.HOS_NAME as hos_name,th.HOS_CONTACT_NAME as hos_contact_name,th.HOS_PHONE_NUMBER as hos_phone_number from t_customer_booking tbk left join t_hospitality th on th.HOS_ID =tbk.hos_id;"
	resp := make([]*TCustomerBookingJoin, 0)

	results, err := db.Query(statement)
	if err != nil {
		log.Print("err:", err)
		return nil, err
	}
	for results.Next() {
		tmpResp := new(TCustomerBookingJoin)
		err = results.Scan(&tmpResp.CustomerBookingId, &tmpResp.CustomerServiceId, &tmpResp.HospitalityId, &tmpResp.StartDate, &tmpResp.EndDate, &tmpResp.HospitalityName, &tmpResp.HospitalityContactName, &tmpResp.HospitalityContactNumber)
		if err != nil {
			log.Print("err:", err)
			return nil, err
		}
		resp = append(resp, tmpResp)
	}

	return resp, nil
}

//GetScholarshipAll
func GetCustomerBookingById(req *apimodelCustomer.CustomerBookingGetRequest) (res []*TCustomerBookingJoin, err error) {

	db := database.GetDB()
	statement := "select tbk.cus_bk_id ,tbk.cus_svc_id ,tbk.hos_id ,tbk.start_date ,tbk.end_date ,th.HOS_NAME as hos_name,th.HOS_CONTACT_NAME as hos_contact_name,th.HOS_PHONE_NUMBER as hos_phone_number from t_customer_booking tbk left join t_hospitality th on th.HOS_ID =tbk.hos_id where CUS_SVC_ID=?;"
	resp := make([]*TCustomerBookingJoin, 0)

	results, err := db.Query(statement, req.CustomerServiceId)
	if err != nil {
		log.Print("err:", err)
		return nil, err
	}
	for results.Next() {
		tmpResp := new(TCustomerBookingJoin)
		err = results.Scan(&tmpResp.CustomerBookingId, &tmpResp.CustomerServiceId, &tmpResp.HospitalityId, &tmpResp.StartDate, &tmpResp.EndDate, &tmpResp.HospitalityName, &tmpResp.HospitalityContactName, &tmpResp.HospitalityContactNumber)
		if err != nil {
			log.Print("err:", err)
			return nil, err
		}
		resp = append(resp, tmpResp)
	}

	return resp, nil
}

//Add New User
func CreateCustomerBooking(req *apimodelCustomer.CustomerBookingCreateRequest, customerBookingId string) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "INSERT INTO hospitality.t_customer_booking (cus_bk_id, cus_svc_id, hos_id, start_date, end_date, CREATE_BY, CREATE_DATE, UPDATE_BY, UPDATE_DATE) VALUES(?, ?, ?, ?,?,?, ?, ?, ?);"
	_, err = tx.Query(statement, customerBookingId, req.CustomerServiceId, req.HospitalityId, req.StartDate, req.EndDate, "SYSTEM", util.TimeNow(), "SYSTEM", util.TimeNow())
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
func UpdateCustomerBooking(req *apimodelCustomer.CustomerBookingUpdateRequest) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "UPDATE hospitality.t_customer_booking SET hos_id=?, start_date=?, end_date=?, UPDATE_DATE=? WHERE cus_bk_id=?;"
	_, err = tx.Query(statement, req.HospitalityId, req.StartDate, req.EndDate, util.TimeNow(), req.CustomerBookingId)
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
func DeleteCustomerBooking(req *apimodelCustomer.CustomerBookingDeleteRequest) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "DELETE FROM hospitality.t_customer_booking WHERE cus_bk_id=?;"
	_, err = tx.Query(statement, req.CustomerBookingId)
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

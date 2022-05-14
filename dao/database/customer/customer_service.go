package databaseCustomer

import (
	"context"
	"log"

	database "github.com/akromibn37/hospitalityCollaboration/dao/database"
	apimodelCustomer "github.com/akromibn37/hospitalityCollaboration/models/customer"
	"github.com/akromibn37/hospitalityCollaboration/pkg/util"
)

type TCustomerService struct {
	CustomerServiceId string  `json:"cus_svc_id"`
	CustomerId        string  `json:"customer_id"`
	ServiceId         *string `json:"svc_id"`
	UserId            *string `json:"user_id"`
	Status            *string `json:"status"`
	CreateBy          string  `json:"create_by"`
	CreateDate        string  `json:"create_date"`
	UpdateBy          string  `json:"update_by"`
	UpdateDate        string  `json:"update_date"`
	StartDate         string  `json:"start_date"`
	EndDate           string  `json:"end_date"`
}

type TCustomerServiceJoin struct {
	CustomerServiceId   string  `json:"cus_svc_id"`
	CustomerId          string  `json:"customer_id"`
	CustomerName        *string `json:"customer_name"`
	CustomerPhoneNumber *string `json:"customer_phone_number"`
	ServiceId           *string `json:"svc_id"`
	ServiceName         *string `json:"svc_name"`
	UserId              *string `json:"user_id"`
	UserName            *string `json:"user_name"`
	Status              *string `json:"status"`
	StartDate           *string `json:"start_date"`
	EndDate             *string `json:"end_date"`
}

//GetScholarshipAll
func GetCustomerServiceAll() (res []*TCustomerServiceJoin, err error) {

	db := database.GetDB()
	statement := "select tcs.cus_svc_id ,tcs.customer_id ,tcs.svc_id ,tcs.user_id ,tcs.status ,tcp.name as customer_name,tcp.phone_number as customer_phone_number,ts.SVC_NAME as svc_name,tup.NAME as user_name,tcs.start_date,tcs.end_date from t_customer_service tcs left join t_customer_profile tcp on tcp.customer_id =tcs.customer_id left join t_user_profile tup on tup.user_id =tcs.user_id left join t_service ts on ts.svc_id = tcs.svc_id;"
	resp := make([]*TCustomerServiceJoin, 0)

	results, err := db.Query(statement)
	if err != nil {
		log.Print("err:", err)
		return nil, err
	}
	for results.Next() {
		tmpResp := new(TCustomerServiceJoin)
		err = results.Scan(&tmpResp.CustomerServiceId, &tmpResp.CustomerId, &tmpResp.ServiceId, &tmpResp.UserId, &tmpResp.Status, &tmpResp.CustomerName, &tmpResp.CustomerPhoneNumber, &tmpResp.ServiceName, &tmpResp.UserName, &tmpResp.StartDate, &tmpResp.EndDate)
		if err != nil {
			log.Print("err:", err)
			return nil, err
		}
		resp = append(resp, tmpResp)
	}

	return resp, nil
}

//Add New User
func CreateCustomerService(customerServiceId string, customerId string, serviceId string) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "INSERT INTO hospitality.t_customer_service (cus_svc_id, customer_id, svc_id, user_id, status, CREATE_BY, CREATE_DATE, UPDATE_BY, UPDATE_DATE) VALUES(?, ?, ?,'','',?, ?, ?, ?);"
	_, err = tx.Query(statement, customerServiceId, customerId, serviceId, "SYSTEM", util.TimeNow(), "SYSTEM", util.TimeNow())
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
func CreateCustomerServiceByStaff(req *apimodelCustomer.CustomerServiceCreateRequest, customerServiceId string) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "INSERT INTO hospitality.t_customer_service (cus_svc_id, customer_id, svc_id, user_id, status, CREATE_BY, CREATE_DATE, UPDATE_BY, UPDATE_DATE,START_DATE,END_DATE) VALUES(?, ?, ?, ?,?,?, ?, ?, ?,?,?);"
	_, err = tx.Query(statement, customerServiceId, req.CustomerId, req.ServiceId, req.UserId, req.Status, "SYSTEM", util.TimeNow(), "SYSTEM", util.TimeNow(), req.StartDate, req.EndDate)
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
func UpdateCustomerService(req *apimodelCustomer.CustomerServiceUpdateRequest) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "UPDATE hospitality.t_customer_service SET customer_id=?, svc_id=?, user_id=?, status=?,start_date=?,end_date=?, UPDATE_DATE=? WHERE cus_svc_id=?;"
	_, err = tx.Query(statement, req.CustomerId, req.ServiceId, req.UserId, req.Status, req.StartDate, req.EndDate, util.TimeNow(), req.CustomerServiceId)
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

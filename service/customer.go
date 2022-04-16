package service

import (
	"time"

	"github.com/akromibn37/hospitalityCollaboration/constant"
	databaseCustomer "github.com/akromibn37/hospitalityCollaboration/dao/database/customer"
	apimodelCustomer "github.com/akromibn37/hospitalityCollaboration/models/customer"
	"github.com/akromibn37/hospitalityCollaboration/pkg/logging"
	projutil "github.com/akromibn37/hospitalityCollaboration/util"
)

type Customer struct {
}

func (c *Customer) CustomerProfileGetAll(t *logging.Timelog) (res *apimodelCustomer.CustomerProfileGetAllResponse, err error) {

	t.TimeInDb = time.Now()
	data, err := databaseCustomer.GetCustomerProfileAll()
	if err != nil {
		return nil, err
	}
	t.TimeOutDb = time.Now()
	resData := make([]apimodelCustomer.CustomerProfileDetail, len(data))
	for i := 0; i < len(data); i++ {
		resData[i].CustomerId = data[i].CustomerId
		resData[i].Name = data[i].Name
		resData[i].LineId = data[i].LineId
		resData[i].PhoneNumber = data[i].PhoneNumber
	}
	res = &apimodelCustomer.CustomerProfileGetAllResponse{Detail: resData}

	return res, nil
}

func (c *Customer) CustomerProfileCreate(req *apimodelCustomer.CustomerProfileCreateRequest, t *logging.Timelog) (res *apimodelCustomer.CustomerProfileCreateResponse, err error) {

	customerId := projutil.GenerateID(constant.CUSTOMER_PROFILE_ID)
	customerServiceId := projutil.GenerateID(constant.CUSTOMER_SERVICE_ID)
	t.TimeInDb = time.Now()
	err = databaseCustomer.CreateCustomerProfile(req, customerId)
	if err != nil {
		return nil, err
	}
	err = databaseCustomer.CreateCustomerService(customerServiceId, customerId, req.ServiceId)
	if err != nil {
		return nil, err
	}
	t.TimeOutDb = time.Now()

	res = &apimodelCustomer.CustomerProfileCreateResponse{}

	return res, nil
}

func (c *Customer) CustomerProfileUpdate(req *apimodelCustomer.CustomerProfileUpdateRequest, t *logging.Timelog) (res *apimodelCustomer.CustomerProfileUpdateResponse, err error) {
	t.TimeInDb = time.Now()
	if err != nil {
		return nil, err
	}
	err = databaseCustomer.UpdateCustomerProfile(req)

	t.TimeOutDb = time.Now()
	res = &apimodelCustomer.CustomerProfileUpdateResponse{}

	return res, nil
}

func (c *Customer) CustomerServiceGetAll(t *logging.Timelog) (res *apimodelCustomer.CustomerServiceGetAllResponse, err error) {

	t.TimeInDb = time.Now()
	data, err := databaseCustomer.GetCustomerServiceAll()
	if err != nil {
		return nil, err
	}
	t.TimeOutDb = time.Now()
	resData := make([]apimodelCustomer.CustomerServiceDetail, len(data))
	for i := 0; i < len(data); i++ {
		resData[i].CustomerServiceId = data[i].CustomerServiceId
		resData[i].CustomerId = data[i].CustomerId
		resData[i].ServiceId = data[i].ServiceId
		resData[i].HospitalityId = data[i].HospitalityId
		resData[i].UserId = data[i].UserId
		resData[i].Status = data[i].Status
		resData[i].CustomerName = data[i].CustomerName
		resData[i].CustomerPhoneNumber = data[i].CustomerPhoneNumber
		resData[i].ServiceName = data[i].ServiceName
		resData[i].HospitalityName = data[i].HospitalityName
		resData[i].HospitalityContactName = data[i].HospitalityContactName
		resData[i].HospitalityContactNumber = data[i].HospitalityContactNumber
		resData[i].UserName = data[i].UserName
	}
	res = &apimodelCustomer.CustomerServiceGetAllResponse{Detail: resData}

	return res, nil
}

func (c *Customer) CustomerServiceUpdate(req *apimodelCustomer.CustomerServiceUpdateRequest, t *logging.Timelog) (res *apimodelCustomer.CustomerServiceUpdateResponse, err error) {
	t.TimeInDb = time.Now()
	if err != nil {
		return nil, err
	}
	err = databaseCustomer.UpdateCustomerService(req)

	t.TimeOutDb = time.Now()
	res = &apimodelCustomer.CustomerServiceUpdateResponse{}

	return res, nil
}

func (c *Customer) CustomerServiceCreate(req *apimodelCustomer.CustomerServiceCreateRequest, t *logging.Timelog) (res *apimodelCustomer.CustomerServiceCreateResponse, err error) {

	customerServiceId := projutil.GenerateID(constant.CUSTOMER_SERVICE_ID)
	t.TimeInDb = time.Now()
	err = databaseCustomer.CreateCustomerServiceByStaff(req, customerServiceId)
	if err != nil {
		return nil, err
	}
	t.TimeOutDb = time.Now()

	res = &apimodelCustomer.CustomerServiceCreateResponse{}

	return res, nil
}

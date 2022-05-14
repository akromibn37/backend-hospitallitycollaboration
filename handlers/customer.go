package handlers

import (
	"net/http"

	"time"

	"github.com/akromibn37/hospitalityCollaboration/constant"
	apimodelCustomer "github.com/akromibn37/hospitalityCollaboration/models/customer"
	"github.com/akromibn37/hospitalityCollaboration/service"
	"github.com/gin-gonic/gin"

	"github.com/akromibn37/hospitalityCollaboration/pkg/app"
	"github.com/akromibn37/hospitalityCollaboration/pkg/e"
	"github.com/akromibn37/hospitalityCollaboration/pkg/logging"
	"github.com/akromibn37/hospitalityCollaboration/pkg/util"
)

// @Summary Get a single article
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /information/get/all [post]
func CustomerProfileGetAll(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_CUSTOMER_PROFILE_GET_ALL
	timeLog.TimeIn = time.Now()

	util.LogRequest(constant.SERVICE_CUSTOMER_PROFILE_GET_ALL, nil)
	//callservice
	getCustomerService := &service.Customer{}
	responseStruct, err := getCustomerService.CustomerProfileGetAll(&timeLog)

	if err != nil {
		util.LogError(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		return
	}

	timeLog.TimeOut = time.Now()
	timeLog.GetDiff()
	util.LogResponse(http.StatusOK, responseStruct)
	util.LogPerformance(timeLog)
	appG.Response(http.StatusOK, e.SUCCESS, responseStruct)
}

// @Summary Get a single article
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /UserCreate [post]
func CustomerProfileCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_CUSTOMER_PROFILE_CREATE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelCustomer.CustomerProfileCreateRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &serviceRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = serviceRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_CUSTOMER_PROFILE_CREATE, serviceRequest)

	//callservice
	getCustomerservice := &service.Customer{}
	responseStruct, err := getCustomerservice.CustomerProfileCreate(&serviceRequest, &timeLog)

	if err != nil {
		// util.LogError(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		appG.ErrResponse(http.StatusOK, "username", err.Error())
		return
	}

	timeLog.TimeOut = time.Now()
	timeLog.GetDiff()
	util.LogResponse(http.StatusOK, responseStruct)
	util.LogPerformance(timeLog)
	appG.Response(http.StatusOK, e.SUCCESS, responseStruct)
}

// @Summary Get a single article
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /update/data [post]
func CustomerProfileUpdate(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_CUSTOMER_PROFILE_UPDATE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelCustomer.CustomerProfileUpdateRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &serviceRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = serviceRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_CUSTOMER_PROFILE_UPDATE, serviceRequest)

	//callservice
	getCustomerservice := &service.Customer{}
	responseStruct, err := getCustomerservice.CustomerProfileUpdate(&serviceRequest, &timeLog)

	if err != nil {
		util.LogError(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		return
	}

	timeLog.TimeOut = time.Now()
	timeLog.GetDiff()
	util.LogResponse(http.StatusOK, responseStruct)
	util.LogPerformance(timeLog)
	appG.Response(http.StatusOK, e.SUCCESS, responseStruct)
}

// @Summary Get a single article
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /information/get/all [post]
func CustomerServiceGetAll(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_CUSTOMER_SERVICE_GET_ALL
	timeLog.TimeIn = time.Now()

	//callservice
	getCustomerService := &service.Customer{}
	responseStruct, err := getCustomerService.CustomerServiceGetAll(&timeLog)

	if err != nil {
		util.LogError(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		return
	}

	timeLog.TimeOut = time.Now()
	timeLog.GetDiff()
	util.LogResponse(http.StatusOK, responseStruct)
	util.LogPerformance(timeLog)
	appG.Response(http.StatusOK, e.SUCCESS, responseStruct)
}

// @Summary Get a single article
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /update/data [post]
func CustomerServiceUpdate(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_CUSTOMER_SERVICE_UPDATE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelCustomer.CustomerServiceUpdateRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &serviceRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = serviceRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_CUSTOMER_SERVICE_UPDATE, serviceRequest)

	//callservice
	getCustomerservice := &service.Customer{}
	responseStruct, err := getCustomerservice.CustomerServiceUpdate(&serviceRequest, &timeLog)

	if err != nil {
		util.LogError(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		return
	}

	timeLog.TimeOut = time.Now()
	timeLog.GetDiff()
	util.LogResponse(http.StatusOK, responseStruct)
	util.LogPerformance(timeLog)
	appG.Response(http.StatusOK, e.SUCCESS, responseStruct)
}

// @Summary Get a single article
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /UserCreate [post]
func CustomerServiceCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_CUSTOMER_SERVICE_CREATE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelCustomer.CustomerServiceCreateRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &serviceRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = serviceRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_CUSTOMER_SERVICE_CREATE, serviceRequest)

	//callservice
	getCustomerservice := &service.Customer{}
	responseStruct, err := getCustomerservice.CustomerServiceCreate(&serviceRequest, &timeLog)

	if err != nil {
		// util.LogError(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		appG.ErrResponse(http.StatusOK, "username", err.Error())
		return
	}

	timeLog.TimeOut = time.Now()
	timeLog.GetDiff()
	util.LogResponse(http.StatusOK, responseStruct)
	util.LogPerformance(timeLog)
	appG.Response(http.StatusOK, e.SUCCESS, responseStruct)
}

// @Summary Get a single article
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /information/get/all [post]
func CustomerBookingGetAll(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_CUSTOMER_BOOKING_GET_ALL
	timeLog.TimeIn = time.Now()

	//callservice
	getCustomerService := &service.Customer{}
	responseStruct, err := getCustomerService.CustomerBookingGetAll(&timeLog)

	if err != nil {
		util.LogError(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		return
	}

	timeLog.TimeOut = time.Now()
	timeLog.GetDiff()
	util.LogResponse(http.StatusOK, responseStruct)
	util.LogPerformance(timeLog)
	appG.Response(http.StatusOK, e.SUCCESS, responseStruct)
}

// @Summary Get a single article
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /update/data [post]
func CustomerBookingUpdate(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_CUSTOMER_BOOKING_UPDATE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelCustomer.CustomerBookingUpdateRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &serviceRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = serviceRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_CUSTOMER_BOOKING_UPDATE, serviceRequest)

	//callservice
	getCustomerservice := &service.Customer{}
	responseStruct, err := getCustomerservice.CustomerBookingUpdate(&serviceRequest, &timeLog)

	if err != nil {
		util.LogError(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		return
	}

	timeLog.TimeOut = time.Now()
	timeLog.GetDiff()
	util.LogResponse(http.StatusOK, responseStruct)
	util.LogPerformance(timeLog)
	appG.Response(http.StatusOK, e.SUCCESS, responseStruct)
}

// @Summary Get a single article
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /UserCreate [post]
func CustomerBookingCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_CUSTOMER_BOOKING_CREATE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelCustomer.CustomerBookingCreateRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &serviceRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = serviceRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_CUSTOMER_BOOKING_CREATE, serviceRequest)

	//callservice
	getCustomerservice := &service.Customer{}
	responseStruct, err := getCustomerservice.CustomerBookingCreate(&serviceRequest, &timeLog)

	if err != nil {
		// util.LogError(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		appG.ErrResponse(http.StatusOK, "username", err.Error())
		return
	}

	timeLog.TimeOut = time.Now()
	timeLog.GetDiff()
	util.LogResponse(http.StatusOK, responseStruct)
	util.LogPerformance(timeLog)
	appG.Response(http.StatusOK, e.SUCCESS, responseStruct)
}

// @Summary Get a single article
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /update/data [post]
func CustomerBookingDelete(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_CUSTOMER_BOOKING_DELETE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelCustomer.CustomerBookingDeleteRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &serviceRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = serviceRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_CUSTOMER_BOOKING_DELETE, serviceRequest)

	//callservice
	getCustomerservice := &service.Customer{}
	responseStruct, err := getCustomerservice.CustomerBookingDelete(&serviceRequest, &timeLog)

	if err != nil {
		util.LogError(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		return
	}

	timeLog.TimeOut = time.Now()
	timeLog.GetDiff()
	util.LogResponse(http.StatusOK, responseStruct)
	util.LogPerformance(timeLog)
	appG.Response(http.StatusOK, e.SUCCESS, responseStruct)
}

// @Summary Get a single article
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /update/data [post]
func CustomerBookingGet(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_CUSTOMER_BOOKING_GET
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelCustomer.CustomerBookingGetRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &serviceRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = serviceRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_CUSTOMER_BOOKING_GET, serviceRequest)

	//callservice
	getCustomerservice := &service.Customer{}
	responseStruct, err := getCustomerservice.CustomerBookingGet(&serviceRequest, &timeLog)

	if err != nil {
		util.LogError(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		return
	}

	timeLog.TimeOut = time.Now()
	timeLog.GetDiff()
	util.LogResponse(http.StatusOK, responseStruct)
	util.LogPerformance(timeLog)
	appG.Response(http.StatusOK, e.SUCCESS, responseStruct)
}

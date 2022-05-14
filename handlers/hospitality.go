package handlers

import (
	"net/http"

	"time"

	"github.com/akromibn37/hospitalityCollaboration/constant"
	apimodelHospitality "github.com/akromibn37/hospitalityCollaboration/models/hospitality"
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
func HospitalityGetAll(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_HOSPITALITY_GET_ALL
	timeLog.TimeIn = time.Now()

	util.LogRequest(constant.SERVICE_HOSPITALITY_GET_ALL, nil)

	//callservice
	getHospitalityService := &service.Hospitality{}
	responseStruct, err := getHospitalityService.HospitalityGetAll(&timeLog)

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
func HospitalityCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_HOSPITALITY_CREATE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelHospitality.HospitalityCreateRequest{}

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
	util.LogRequest(constant.SERVICE_HOSPITALITY_CREATE, serviceRequest)

	//callservice
	getHospitalityservice := &service.Hospitality{}
	responseStruct, err := getHospitalityservice.HospitalityCreate(&serviceRequest, &timeLog)

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
func HospitalityUpdate(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_HOSPITALITY_UPDATE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelHospitality.HospitalityUpdateRequest{}

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
	util.LogRequest(constant.SERVICE_HOSPITALITY_UPDATE, serviceRequest)

	//callservice
	getHospitalityservice := &service.Hospitality{}
	responseStruct, err := getHospitalityservice.HospitalityUpdate(&serviceRequest, &timeLog)

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
func HospitalityTypeGetAll(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_HOSPITALITY_TYPE_GET_ALL
	timeLog.TimeIn = time.Now()

	util.LogRequest(constant.SERVICE_HOSPITALITY_TYPE_GET_ALL, nil)

	//callservice
	getHospitalityService := &service.Hospitality{}
	responseStruct, err := getHospitalityService.HospitalityTypeGetAll(&timeLog)

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
func HospitalityTypeCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_HOSPITALITY_TYPE_CREATE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelHospitality.HospitalityTypeCreateRequest{}

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
	util.LogRequest(constant.SERVICE_HOSPITALITY_TYPE_CREATE, serviceRequest)

	//callservice
	getHospitalityservice := &service.Hospitality{}
	responseStruct, err := getHospitalityservice.HospitalityTypeCreate(&serviceRequest, &timeLog)

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
func HospitalityTypeUpdate(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_HOSPITALITY_TYPE_UPDATE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelHospitality.HospitalityTypeUpdateRequest{}

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
	util.LogRequest(constant.SERVICE_HOSPITALITY_TYPE_UPDATE, serviceRequest)

	//callservice
	getHospitalityservice := &service.Hospitality{}
	responseStruct, err := getHospitalityservice.HospitalityTypeUpdate(&serviceRequest, &timeLog)

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
func HospitalityTypeDelete(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_HOSPITALITY_TYPE_DELETE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelHospitality.HospitalityTypeDeleteRequest{}

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
	util.LogRequest(constant.SERVICE_HOSPITALITY_TYPE_DELETE, serviceRequest)

	//callservice
	getHospitalityservice := &service.Hospitality{}
	responseStruct, err := getHospitalityservice.HospitalityTypeDelete(&serviceRequest, &timeLog)

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

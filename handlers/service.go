package handlers

import (
	"net/http"

	"time"

	"github.com/akromibn37/hospitalityCollaboration/constant"
	apimodelService "github.com/akromibn37/hospitalityCollaboration/models/service"
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
func ServiceGetAll(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_SERVICE_GET_ALL
	timeLog.TimeIn = time.Now()

	util.LogRequest(constant.SERVICE_SERVICE_GET_ALL, nil)
	//callservice
	getServiceService := &service.Service{}
	responseStruct, err := getServiceService.ServiceGetAll(&timeLog)

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
func ServiceCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_SERVICE_CREATE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelService.ServiceCreateRequest{}

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
	util.LogRequest(constant.SERVICE_SERVICE_CREATE, serviceRequest)

	//callservice
	getServiceservice := &service.Service{}
	responseStruct, err := getServiceservice.ServiceCreate(&serviceRequest, &timeLog)

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
func ServiceUpdate(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_SERVICE_UPDATE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelService.ServiceUpdateRequest{}

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
	util.LogRequest(constant.SERVICE_SERVICE_UPDATE, serviceRequest)

	//callservice
	getServiceservice := &service.Service{}
	responseStruct, err := getServiceservice.ServiceUpdate(&serviceRequest, &timeLog)

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
func ServiceDelete(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_SERVICE_DELETE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	serviceRequest := apimodelService.ServiceDeleteRequest{}

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
	util.LogRequest(constant.SERVICE_SERVICE_DELETE, serviceRequest)

	//callservice
	getServiceservice := &service.Service{}
	responseStruct, err := getServiceservice.ServiceDelete(&serviceRequest, &timeLog)

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

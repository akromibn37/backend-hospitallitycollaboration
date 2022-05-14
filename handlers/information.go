package handlers

import (
	"net/http"

	"time"

	"github.com/akromibn37/hospitalityCollaboration/constant"
	apimodelInformation "github.com/akromibn37/hospitalityCollaboration/models/information"
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
// @Router /get/profile [post]
func InformationGetProfile(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_INFORMATION_GET_PROFILE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	informationRequest := apimodelInformation.InformationGetProfileRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &informationRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = informationRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_INFORMATION_GET_PROFILE, informationRequest)

	//callservice
	getInformationService := &service.Information{}
	responseStruct, err := getInformationService.InformationGetProfile(&informationRequest, &timeLog)

	if err != nil {
		util.LogError(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		appG.Response(http.StatusOK, e.SUCCESS, "test")
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
// @Router /get/profile [post]
func InformationGetAllProfile(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_INFORMATION_GET_ALL_PROFILE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	informationRequest := apimodelInformation.InformationGetAllProfileRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &informationRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	// if err = informationRequest.Validate(); nil != err {
	// 	app.ValidError(err)
	// 	appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
	// 	return
	// }
	util.LogRequest(constant.SERVICE_INFORMATION_GET_ALL_PROFILE, informationRequest)

	//callservice
	getInformationService := &service.Information{}
	responseStruct, err := getInformationService.InformationGetAllProfile(&timeLog)

	if err != nil {
		util.LogError(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
		appG.Response(http.StatusOK, e.SUCCESS, "test")
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
// @Router /update/profile [post]
func InformationUpdateProfile(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_INFORMATION_UPDATE_PROFILE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	informationRequest := apimodelInformation.InformationUpdateProfileRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &informationRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = informationRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_INFORMATION_UPDATE_PROFILE, informationRequest)

	//callservice
	getInformationService := &service.Information{}
	responseStruct, err := getInformationService.InformationUpdateProfile(&informationRequest, &timeLog)

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

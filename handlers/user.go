package handlers

import (
	"net/http"

	"time"

	"github.com/akromibn37/hospitalityCollaboration/constant"
	apimodel "github.com/akromibn37/hospitalityCollaboration/models/user"
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
// @Router /UserCreate [post]
func UserCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_USER_CREATE
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	userRequest := apimodel.UserCreateRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &userRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = userRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_USER_CREATE, userRequest)

	//callservice
	getUserservice := &service.User{}
	responseStruct, err := getUserservice.UserCreate(&userRequest, &timeLog)

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
// @Router /UserCreate [post]
func UserCreateByAdmin(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_USER_CREATE_BY_ADMIN
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	userRequest := apimodel.UserCreateByAdminRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &userRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = userRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_USER_CREATE_BY_ADMIN, userRequest)

	//callservice
	getUserservice := &service.User{}
	responseStruct, err := getUserservice.UserCreateByAdmin(&userRequest, &timeLog)

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
// @Router /UserLogin [post]
func UserLogin(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_USER_LOGIN
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	userRequest := apimodel.UserLoginRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &userRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = userRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_USER_LOGIN, userRequest)

	//callservice
	getUserservice := &service.User{}
	responseStruct, err := getUserservice.UserLogin(&userRequest, &timeLog)

	if err != nil {
		util.LogError(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
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
// @Router /UserLogin [post]
func UserGetByUserId(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_USER_GET_BY_USERID
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	userRequest := apimodel.UserGetByUserIdRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &userRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = userRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_USER_GET_BY_USERID, userRequest)

	//callservice
	getUserservice := &service.User{}
	responseStruct, err := getUserservice.UserGetByUserId(&userRequest, &timeLog)

	if err != nil {
		util.LogError(http.StatusInternalServerError, e.ERROR_CREATE_USER, err.Error())
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
// @Router /update/password [post]
func UserUpdatePassword(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_USER_UPDATE_PASSWORD
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	userRequest := apimodel.UserUpdatePasswordRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &userRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = userRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_USER_UPDATE_PASSWORD, userRequest)

	//callservice
	getUserservice := &service.User{}
	responseStruct, err := getUserservice.UserUpdatePassword(&userRequest, &timeLog)

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
func UserUpdateData(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_USER_UPDATE_DATA
	timeLog.TimeIn = time.Now()

	request, err := c.GetRawData()
	userRequest := apimodel.UserUpdateDataRequest{}

	if err != nil {
		util.LogError(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.GET_REQUEST_FAIL, err.Error())
		return
	}
	if err = util.UnpackRequest(request, &userRequest); nil != err {
		util.LogError(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		appG.Response(http.StatusBadRequest, e.JSON_UNMARSHAL_FAIL, err.Error())
		return
	}

	//validate request
	if err = userRequest.Validate(); nil != err {
		app.ValidError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	util.LogRequest(constant.SERVICE_USER_UPDATE_DATA, userRequest)

	//callservice
	getUserservice := &service.User{}
	responseStruct, err := getUserservice.UserUpdateData(&userRequest, &timeLog)

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
func UserGetAll(c *gin.Context) {
	appG := app.Gin{C: c}
	timeLog := logging.Timelog{}
	timeLog.FuncNm = constant.SERVICE_USER_GET_ALL_DATA
	timeLog.TimeIn = time.Now()

	util.LogRequest(constant.SERVICE_USER_GET_ALL_DATA, nil)
	//callservice
	getUserservice := &service.User{}
	responseStruct, err := getUserservice.UserGetAll(&timeLog)

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

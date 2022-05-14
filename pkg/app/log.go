package app

import (
	"github.com/akromibn37/hospitalityCollaboration/pkg/logging"
	"github.com/astaxie/beego/validation"
)

// MarkErrors logs validation error
func MarkErrors(errors []*validation.Error) {

	for _, err := range errors {
		logging.Error(err.Key, err.Message)
	}
	return
}

func ValidError(errors error) {
	logging.Error(errors)
}

//logs error logs
func LogError(httpCode int, errMsg string, msg string) {
	logging.Error(httpCode, errMsg, msg)
	return
}

//LogInterface logs interface logs request
func LogInterfaceRequest(serviceName string, data interface{}) {
	logging.Info("Request: ", serviceName, data)
	return
}

//LogInterface logs interface logs request
func LogInterfaceInformation(serviceName string, data interface{}) {
	logging.Info("Information: ", serviceName, data)
	return
}

// LogInterface logs interface logs response
func LogInterfaceResponse(httpCode int, data interface{}) {
	logging.Info("Response: ", httpCode, data)
	return
}

//LogPerformance logs performance logs
func LogPerformance(t logging.Timelog) {
	logging.PerfInfo("Performance:", t.FuncNm, "E2E=", t.DiffE2e, "DBtime=", t.DiffDB)
}

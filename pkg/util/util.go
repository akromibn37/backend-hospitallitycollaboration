package util

import (
	"fmt"
	"time"

	"github.com/akromibn37/hospitalityCollaboration/pkg/app"
	"github.com/akromibn37/hospitalityCollaboration/pkg/e"
	"github.com/akromibn37/hospitalityCollaboration/pkg/logging"
	"github.com/akromibn37/hospitalityCollaboration/pkg/setting"
)

func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}

func TimeNow() string {
	t := time.Now()
	return t.Format(e.TIME_FORMAT_ISO)
}

func LogError(httpCode int, errCode int, msg string) {
	errMsg := e.GetMsg(errCode)
	app.LogError(httpCode, errMsg, msg)
	return
}

func LogRequest(data interface{}) {
	req := fmt.Sprintf("%#v", data)
	app.LogInterfaceRequest(req)
	return
}

func LogResponse(httpCode int, data interface{}) {
	res := fmt.Sprintf("%#v", data)
	app.LogInterfaceResponse(httpCode, res)
	return
}

func LogPerformance(timedata logging.Timelog) {
	app.LogPerformance(timedata)
	return
}

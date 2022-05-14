package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akromibn37/hospitalityCollaboration/dao/database"
	"github.com/akromibn37/hospitalityCollaboration/pkg/logging"
	"github.com/akromibn37/hospitalityCollaboration/pkg/setting"
	"github.com/akromibn37/hospitalityCollaboration/pkg/util"
	"github.com/akromibn37/hospitalityCollaboration/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	database.Setup()
	// redis.Setup()
	logging.Setup()
	util.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/EDDYCJY/go-gin-example
// @license.name MIT
// @license.url https://github.com/EDDYCJY/go-gin-example/blob/master/LICENSE
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouters()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}

package app

import (
	"github.com/akromibn37/hospitalityCollaboration/pkg/e"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Error []Error `json:"error"`
}
type Error struct {
	ErrType string `json:"type"`
	Msg     string `json:"message"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	g.C.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	g.C.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	g.C.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

// Response setting gin.JSON
func (g *Gin) ErrResponse(httpCode int, errType string, errMsg string) {
	g.C.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	g.C.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	g.C.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	g.C.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	g.C.JSON(httpCode, ErrorResponse{
		Error: []Error{
			{
				ErrType: errType,
				Msg:     errMsg,
			},
		},
	})
	return
}

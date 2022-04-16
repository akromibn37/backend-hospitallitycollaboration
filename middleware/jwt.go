package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/akromibn37/hospitalityCollaboration/pkg/e"
	"github.com/akromibn37/hospitalityCollaboration/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		bearToken := c.Request.Header.Get("Authorization")
		// token := c.Query("token")
		strArr := strings.Split(bearToken, " ")
		var token string
		if len(strArr) == 2 {
			token = strArr[1]
		}
		// log.Print(token)
		if token == "" {
			log.Print("here")
			code = e.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

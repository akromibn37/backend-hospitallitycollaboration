package routers

import (
	"net/http"

	"github.com/akromibn37/hospitalityCollaboration/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORS)
	r.Static("images", "../images")
	//admin section
	user := r.Group("/api/v1/user")
	service := r.Group("/api/v1/service")
	hospitality := r.Group("/api/v1/hospitality")
	information := r.Group("/api/v1/information")
	customer := r.Group("/api/v1/customer")
	// apiv1.Use(middleware.JWT())
	user.Use()
	{
		user.POST("/login", handlers.UserLogin)
		user.POST("/get/userId", handlers.UserGetByUserId)
		user.POST("/create", handlers.UserCreateByAdmin)
		user.GET("/get/all", handlers.UserGetAll)
		user.POST("/update/password", handlers.UserUpdatePassword)
		user.POST("/update/data", handlers.UserUpdateData)
	}

	service.Use()
	{
		service.GET("/get/all", handlers.ServiceGetAll)
		service.POST("/create", handlers.ServiceCreate)
		service.POST("/update", handlers.ServiceUpdate)
	}

	information.Use()
	{
		information.POST("/get/profile", handlers.InformationGetProfile)
		information.POST("/get/profile/all", handlers.InformationGetAllProfile)
		information.POST("/update/profile", handlers.InformationUpdateProfile)
		// information.POST("/update/family", handlers.UserUpdateData)
		// information.POST("/update/document", handlers.UserUpdateData)
	}

	hospitality.Use()
	{
		hospitality.GET("/get/all", handlers.HospitalityGetAll)
		hospitality.POST("/create", handlers.HospitalityCreate)
		hospitality.POST("/update", handlers.HospitalityUpdate)
	}

	customer.Use()
	{
		customer.GET("/profile/get/all", handlers.CustomerProfileGetAll)
		customer.POST("/profile/update", handlers.CustomerProfileUpdate)
		customer.GET("/service/get/all", handlers.CustomerServiceGetAll)
		customer.POST("/service/update", handlers.CustomerServiceUpdate)
		customer.POST("/service/create", handlers.CustomerServiceCreate)
	}

	//create user on table t_user_master,t_user_common
	r.POST("/user/create", handlers.UserCreate)
	r.POST("/customer/profile/create", handlers.CustomerProfileCreate)

	return r

}

// CORS Middleware
func CORS(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}

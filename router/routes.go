package router

import (
	"github.com/gin-gonic/gin"
	"github.com/steixeira93/medical-api/handler"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	// create routes to handle with medical schedule
	{
		v1.POST("/appointment", handler.CreateAppointment)
	}
}
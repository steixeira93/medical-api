package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	// create routes to handle with medical schedule
	{
	}
}
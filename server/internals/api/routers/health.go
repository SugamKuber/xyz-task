package routers

import (
	"github.com/gin-gonic/gin"
	"server/internals/api/controllers"
)

func HealthRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/h", controllers.Health)
	}
}

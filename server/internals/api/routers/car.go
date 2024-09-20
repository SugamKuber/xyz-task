package routers

import (
	"github.com/gin-gonic/gin"
	"server/internals/api/controllers"
)

func CarRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/upload", controllers.UploadCarImage)
	}
}

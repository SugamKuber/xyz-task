package routers

import (
    "server/internals/api/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    apiGroup := router.Group("/api")
    {
        apiGroup.GET("/h", controllers.Health)
    }
}

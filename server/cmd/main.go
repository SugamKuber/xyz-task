package main

import (
	"github.com/gin-gonic/gin"
	"server/internals/api/routers"
	db "server/internals/init"
	"server/internals/middlewares"
)

func main() {
	db.InitDB()

	router := gin.Default()
	router.Use(middlewares.CORSConfig())
	router.Use(middlewares.Logger())
	routers.HealthRoutes(router)
	routers.CarRoutes(router)

	router.Run(":8080")
}

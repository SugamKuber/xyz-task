package main

import (
    "server/internals/api/routers"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    routers.SetupRoutes(router)
    router.Run(":8080")
}

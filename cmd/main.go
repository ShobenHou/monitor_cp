package main

import (
	"github.com/ShobenHou/monitor_cp/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api.SetupRoutes(router)

	router.Run()
}

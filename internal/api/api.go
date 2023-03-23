package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/agents/:id/config", GetAgentConfig)
	router.PUT("/agents/:id/config", UpdateAgentConfig)
}

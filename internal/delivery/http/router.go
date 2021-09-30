package http

import (
	"github.com/gin-gonic/gin"
)

func RouterInit(engine *gin.Engine, handlers *Handler) {

	group := engine.Group("/auth")

	group.GET("/health-check", handlers.HealthCheck)
	group.POST("/v1/register", handlers.Register)
	//group.POST("/v2/login", handlers.Login)
}


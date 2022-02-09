package http

import (
	"github.com/gin-gonic/gin"
)

func RouterInit(engine *gin.Engine, handlers *Handler) {

	group := engine.Group("/auth")

	group.GET("/health-check", handlers.HealthCheck)
	group.POST("/user/v1/register", handlers.Register)
	group.POST("/user/v1/login", handlers.Login)
	group.PUT("/user/v1", AuthenticationMiddleware(handlers.Update))
	group.DELETE("/user/v1", AuthenticationMiddleware(handlers.Delete))
}

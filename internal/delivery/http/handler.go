package http

import (
	"github.com/felipeagger/go-boilerplate/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

// HealthCheck godoc
// @Summary Return status of service
// @Tags General
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /health-check [get]
func (h *Handler) HealthCheck(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Ok",
	})
}

// Register godoc
// @Summary Endpoint to signup user
// @Description Endpoint to register user
// @Tags Register
// @Accept  json
// @Produce  json
// @Param X-Client-Id header string true "Client identifier"
// @Param X-Authorization header string true "Auth Token"
// @Param Payload body domain.Signup true "Payload"
// @Success 200 {object} domain.Signup
// @Failure 400 {object} string
// @Router /v1 [post]
func (h *Handler) Register(c *gin.Context) {

	clientId := c.GetHeader("X-Client-Id")
	if clientId == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, "X-Client-Id not found in headers")
		return
	}

	token := c.GetHeader("X-Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, "X-Authorization not found in headers")
		return
	}

	var payload domain.Signup
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Ok")
}

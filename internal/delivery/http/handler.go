package http

import (
	"github.com/felipeagger/go-boilerplate/internal/config"
	"github.com/felipeagger/go-boilerplate/pkg/trace"
	"github.com/felipeagger/go-boilerplate/pkg/utils"
	"net/http"
	"strconv"

	"github.com/felipeagger/go-boilerplate/internal/controller"
	"github.com/felipeagger/go-boilerplate/internal/domain"
	"github.com/gin-gonic/gin"
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
// @Param Payload body domain.Signup true "Payload"
// @Success 201 {object} domain.Signup
// @Failure 400 {object} string
// @Router /user/v1/register [post]
func (h *Handler) Register(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.Register")
	defer span.End()

	clientID := c.GetHeader("X-Client-Id")
	if clientID == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, "X-Client-Id not found in headers")
		return
	}

	trace.AddSpanTags(span, map[string]string{"app.client_id": clientID})

	var payload domain.Signup
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Unprocessable Entity")
		return
	}

	err := controller.CreateUser(ctx, payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "internal error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "Created"})
}

// Login godoc
// @Summary Endpoint to login user
// @Description Endpoint to login user
// @Tags Login
// @Accept  json
// @Produce  json
// @Param X-Client-Id header string true "Client identifier"
// @Param Payload body domain.Login true "Payload"
// @Success 200 {object} domain.LoginResponse
// @Failure 401 {object} domain.LoginResponse
// @Router /user/v1/login [post]
func (h *Handler) Login(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.Login")
	defer span.End()

	clientID := c.GetHeader("X-Client-Id")
	if clientID == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, "X-Client-Id not found in headers")
		return
	}

	trace.AddSpanTags(span, map[string]string{"app.client_id": clientID})

	var payload domain.Login
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Unprocessable Entity")
		return
	}

	response, err := controller.SignInUser(ctx, payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Status Unauthorized")
		return
	}

	c.JSON(http.StatusOK, response)
}

// Update godoc
// @Summary Endpoint to update user
// @Description Endpoint to update user
// @Tags Update
// @Accept  json
// @Produce  json
// @Param X-Client-Id header string true "Client identifier"
// @Param X-Authorization header string true "Auth Token"
// @Param Payload body domain.Signup true "Payload"
// @Success 200 {object} domain.Signup
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/v1 [put]
func (h *Handler) Update(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.Update")
	defer span.End()

	clientID := c.GetHeader("X-Client-Id")
	if clientID == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, "X-Client-Id not found in headers")
		return
	}

	trace.AddSpanTags(span, map[string]string{"app.client_id": clientID})

	token := c.GetHeader("X-Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, "X-Authorization not found in headers")
		trace.FailSpan(span, "Unauthorized")
		return
	}

	userID, err := utils.ValidateJWT(token, config.GetEnv().TokenSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Unauthorized")
		return
	}

	var payload domain.Signup
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Unprocessable Entity")
		return
	}

	ID, _ := strconv.ParseInt(userID, 10, 64)
	if err := controller.UpdateUser(ctx, ID, payload); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Updated"})
}

package http

import (
	"github.com/felipeagger/go-boilerplate/internal/entity"
	"github.com/felipeagger/go-boilerplate/internal/usecase/user"
	"github.com/felipeagger/go-boilerplate/pkg/trace"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	UserSvc *user.Service
}

func NewHandler(userSvc *user.Service) *Handler {
	return &Handler{
		UserSvc: userSvc,
	}
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
// @Param Payload body entity.Signup true "Payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} string
// @Router /user/v1/register [post]
func (h *Handler) Register(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.Register")
	defer span.End()

	clientID, abort := commonsValidations(c)
	if abort {
		return
	}

	trace.AddSpanTags(span, map[string]string{"app.client_id": clientID})

	var payload entity.Signup
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Unprocessable Entity")
		return
	}

	err := h.UserSvc.CreateUser(ctx, payload)
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
// @Param Payload body entity.Login true "Payload"
// @Success 200 {object} entity.LoginResponse
// @Failure 401 {object} entity.LoginResponse
// @Router /user/v1/login [post]
func (h *Handler) Login(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.Login")
	defer span.End()

	clientID, abort := commonsValidations(c)
	if abort {
		return
	}

	trace.AddSpanTags(span, map[string]string{"app.client_id": clientID})

	var payload entity.Login
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Unprocessable Entity")
		return
	}

	response, err := h.UserSvc.SignInUser(ctx, payload)
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
// @Param Payload body entity.Signup true "Payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/v1 [put]
func (h *Handler) Update(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.Update")
	defer span.End()

	clientID, abort := commonsValidations(c)
	if abort {
		return
	}

	trace.AddSpanTags(span, map[string]string{"app.client_id": clientID, "app.user_id": c.GetHeader("X-User-Id")})
	userID, _ := strconv.ParseInt(c.GetHeader("X-User-Id"), 10, 64)

	var payload entity.Signup
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Unprocessable Entity")
		return
	}

	if err := h.UserSvc.UpdateUser(ctx, userID, payload); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Updated"})
}

// Delete godoc
// @Summary Endpoint to delete user
// @Description Endpoint to delete user
// @Tags Delete
// @Accept  json
// @Produce  json
// @Param X-Client-Id header string true "Client identifier"
// @Param X-Authorization header string true "Auth Token"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/v1 [delete]
func (h *Handler) Delete(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.Delete")
	defer span.End()

	clientID, abort := commonsValidations(c)
	if abort {
		return
	}

	trace.AddSpanTags(span, map[string]string{"app.client_id": clientID, "app.user_id": c.GetHeader("X-User-Id")})
	userID, _ := strconv.ParseInt(c.GetHeader("X-User-Id"), 10, 64)

	if err := h.UserSvc.DeleteUser(ctx, userID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Bad Request")
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Deleted"})
}
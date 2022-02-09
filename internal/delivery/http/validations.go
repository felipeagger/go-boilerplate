package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func commonsValidations(c *gin.Context) (string, bool) {

	clientID := c.GetHeader("X-Client-Id")
	if clientID == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, "X-Client-Id not found in headers")
		return "", true
	}

	return clientID, false
}

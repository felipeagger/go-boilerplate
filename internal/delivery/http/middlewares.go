package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/felipeagger/go-boilerplate/internal/config"
	"github.com/felipeagger/go-boilerplate/pkg/cache"
	"github.com/felipeagger/go-boilerplate/pkg/trace"
	"github.com/felipeagger/go-boilerplate/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func validateSession(ctx context.Context, token string) (string, error) {

	userID, err := utils.ValidateJWT(token, config.GetEnv().TokenSecret)
	if err != nil {
		return "", err
	}

	tkn, err := cache.GetCacheClient().Get(ctx, fmt.Sprintf("tkn-%s", userID))
	if err != nil {
		return "", err
	}

	if tkn != token {
		return "", errors.New(utils.ErrorInvalidToken)
	}

	return userID, nil
}

func AuthenticationMiddleware(function func(c *gin.Context)) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, span := trace.NewSpan(c.Request.Context(), "Auth.Middleware")
		defer span.End()

		token := c.GetHeader("X-Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.ErrorAuthHeaderNotFound)
			trace.AddSpanError(span, errors.New(utils.ErrorAuthHeaderNotFound))
			trace.FailSpan(span, utils.ErrorUnauthorized)
			return
		}

		userID, err := validateSession(ctx, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			trace.AddSpanError(span, err)
			trace.FailSpan(span, utils.ErrorUnauthorized)
			return
		}

		c.Request.Header.Set("X-User-Id", userID)
		function(c)
	}
}
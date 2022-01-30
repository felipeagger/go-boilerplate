package utils

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LoggerInterface interface {
	Debug(i ...interface{})
	Info(i ...interface{})
	Error(i ...interface{})
}

type Logger struct {
}

func (l *Logger) Debug(i ...interface{}) {
	log.Printf("[DEBUG] %s", fmt.Sprintln(i...))
}

func (l *Logger) Info(i ...interface{}) {
	log.Printf("[INFO] %s", fmt.Sprintln(i...))
}

func (l *Logger) Error(i ...interface{}) {
	log.Printf("[ERROR] %s", fmt.Sprintln(i...))
}

// Logger is the logrus logger handler
func LogMiddleware(logger logrus.FieldLogger) gin.HandlerFunc {

	return func(c *gin.Context) {

		if !strings.Contains(c.Request.URL.Path, "/health-check") {

			c.Next()
			statusCode := c.Writer.Status()

			entry := logger.WithFields(logrus.Fields{
				"log_version": "1.0",
				"request_id":  "undefined",
				"date_time":   time.Now(),
				"product": map[string]interface{}{
					"name":        "go-boilerplate",
					"application": "go-boilerplate",
					"version":     "0.0.1",
					"channel":     c.GetHeader("X-Client-Id"),
					"http": map[string]string{
						"method": c.Request.Method,
						"path":   c.Request.URL.Path,
					},
				},
				"user": map[string]string{
					"uuid":  c.GetHeader("userId"),
					"token": c.GetHeader("x-bifrost-authorization"),
				},
				"origin": map[string]interface{}{
					"client":      c.GetHeader("X-Client-Id"),
					"application": "go-boilerplate",
					"ip":          c.ClientIP(),
					"headers": map[string]string{
						"user_agent": c.Request.UserAgent(),
						"origin":     c.GetHeader("Origin"),
						"referer":    c.Request.Referer(),
					},
				},
				"context": map[string]interface{}{
					"service":     "",
					"status_code": c.Writer.Status(),
					"data": map[string]string{
						"message": "",
						"content": "",
					},
				},
			})

			if len(c.Errors) > 0 {
				entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
			} else {
				if statusCode > 499 {
					entry.Error()
				} else if statusCode > 399 {
					entry.Warn()
				} else {
					entry.Info()
				}
			}
		}
	}
}

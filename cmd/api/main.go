package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/felipeagger/go-boilerplate/internal/config"
	httpd "github.com/felipeagger/go-boilerplate/internal/delivery/http"
	"github.com/felipeagger/go-boilerplate/pkg/utils"

	_ "github.com/felipeagger/go-boilerplate/docs"
)

// @title Go Boilerplate API
// @version 1.0
// @description This is Example API in Go.

// @host localhost:8000
// @BasePath /auth
// @schemes http
// @query.collection.format multi

// @x-extension-openapi {"example": "value on a json format"}

func main() {
	ctx := context.Background()

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	// Bootstrap tracer.
	//prv, err := trace.NewProvider(trace.ProviderConfig{
	//	JaegerEndpoint: fmt.Sprintf("http://%s/api/traces", config.GetEnv().TraceHost),
	//	ServiceName:    "client",
	//	ServiceVersion: "1.0.0",
	//	Environment:    "dev",
	//	Disabled:       false,
	//})

	//if err != nil {
	//	log.Fatalln(err)
	//}

	//defer prv.Close(ctx)

	engine := gin.New()
	engine.Use(cors.Default(),
		utils.LogMiddleware(log),
		gin.Recovery(),
		gzip.Gzip(gzip.DefaultCompression))

	handler := httpd.NewHandler()
	httpd.RouterInit(engine, &handler)

	engine.GET("/auth/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	gin.SetMode(gin.ReleaseMode)
	if config.GetEnv().Debug {
		gin.SetMode(gin.DebugMode)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.GetEnv().Port),
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down API...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("API Server forced to shutdown:", err)
	}

	log.Println("API Server exiting")
}

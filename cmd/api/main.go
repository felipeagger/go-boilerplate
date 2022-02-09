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

	httpd "github.com/felipeagger/go-boilerplate/internal/delivery/http"
	"github.com/felipeagger/go-boilerplate/internal/repository"
	"github.com/felipeagger/go-boilerplate/internal/usecase/user"
	"github.com/felipeagger/go-boilerplate/pkg/cache"
	"github.com/felipeagger/go-boilerplate/pkg/database"

	"github.com/felipeagger/go-boilerplate/pkg/trace"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/felipeagger/go-boilerplate/internal/config"
	"github.com/felipeagger/go-boilerplate/pkg/utils"

	_ "github.com/felipeagger/go-boilerplate/docs"
)

// @title Go Boilerplate API
// @version 1.0
// @description This is Example API in Go.

// @host 0.0.0.0:8000
// @BasePath /auth
// @schemes http
// @query.collection.format multi

func init() {
	cache.InitCacheClientSvc(config.GetEnv().CacheHost, config.GetEnv().CachePort, config.GetEnv().CachePassword)
}

func main() {
	ctx := context.Background()

	dbInstance, err := database.NewMySQLConnection(config.GetEnv().DBHost, config.GetEnv().DBName,
		config.GetEnv().DBUser, config.GetEnv().DBPass)
	if err != nil {
		fmt.Println(utils.ErrorDatabaseConn)
		panic(err)
	}

	err = repository.Migrate(dbInstance)
	if err != nil {
		fmt.Println(utils.ErrorDatabaseMigrate)
		panic(err)
	}

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	// Bootstrap tracer.
	prv, err := trace.NewProvider(trace.ProviderConfig{
		JaegerEndpoint: fmt.Sprintf("http://%s/api/traces", config.GetEnv().TraceHost),
		ServiceName:    config.ServiceName,
		ServiceVersion: "1.0.0",
		Environment:    config.GetEnv().Env,
		Disabled:       false,
	})

	if err != nil {
		log.Fatalln(err)
	}

	defer prv.Close(ctx)

	engine := gin.New()

	customCors := cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	})

	engine.Use(
		customCors,
		utils.LogMiddleware(log),
		otelgin.Middleware(config.ServiceName),
		gin.Recovery(),
		gzip.Gzip(gzip.DefaultCompression))

	userRepository := repository.NewGORMUserRepository(dbInstance)
	userService := user.NewService(userRepository, cache.GetCacheClient())

	handler := httpd.NewHandler(userService)
	httpd.RouterInit(engine, handler)

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

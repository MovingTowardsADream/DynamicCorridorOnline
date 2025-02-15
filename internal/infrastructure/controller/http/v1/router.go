package v1

import (
	"log/slog"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	authroutes "TicTacToe/internal/infrastructure/controller/http/v1/auth"
)

func NewRouter(log *slog.Logger, handler *gin.Engine, auth authroutes.EditInfo) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	authHandler := handler.Group("/auth")
	{
		authroutes.NewAuthRoutes(log, authHandler, auth)
	}
	//h := handler.Group("/api/v1")
	//{
	//
	//}
}

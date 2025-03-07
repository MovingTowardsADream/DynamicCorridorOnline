package v1

import (
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"TicTacToe/internal/infrastructure/config"
	"TicTacToe/internal/interfaces/middleware"
)

func NewRouter(cfg *config.Config, handler *gin.Engine, user EditInfo, stat StatistInfo) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Middleware
	mw := middleware.New(cfg.Frontend.Host, user)

	handler.Use(mw.CORSMiddleware.CORSMiddleware())

	// Routers
	authHandler := handler.Group("/user")
	{
		newAuthRoutes(authHandler, user)
	}
	apiHandler := handler.Group("/api/v1", mw.UserIdentity())
	{
		newStatistRoutes(apiHandler, stat)
		newGameRoutes(apiHandler)
	}
}

package handlers

import (
	"net/http"
	"time"

	"github.com/Taehoya/pocket-mate/pkg/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	TripUseCase TripUseCase
}

func New(TripUseCase TripUseCase) *Handler {
	return &Handler{
		TripUseCase: TripUseCase,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	engine := gin.Default()

	engine.Use(middlewares.LoggingMiddleware())
	engine.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := engine.Group("api/v1")
	router.GET("healthcheck", healthCheck())

	return engine
}

func healthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"time":    time.Now().Format(time.RFC3339),
		})
	}
}

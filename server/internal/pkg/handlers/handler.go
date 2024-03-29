package handlers

import (
	"net/http"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	TripUseCase        TripUseCase
	CountryUsecase     CountryUsecase
	UserUseCase        UserUseCase
	TransactionUseCase TransactionUseCase
}

func New(TripUseCase TripUseCase, CountryUseCase CountryUsecase, UserUseCase UserUseCase, TransactionUseCase TransactionUseCase) *Handler {
	return &Handler{
		TripUseCase:        TripUseCase,
		CountryUsecase:     CountryUseCase,
		UserUseCase:        UserUseCase,
		TransactionUseCase: TransactionUseCase,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	engine := gin.Default()
	engine.Use(middlewares.LoggingMiddleware())
	engine.GET("/", healthCheck())

	apiGroup := engine.Group("api/v1")
	apiGroup.GET("/healthcheck", healthCheck())
	apiGroup.GET("/countries", h.GetCountries)
	apiGroup.POST("/users", h.Register)
	apiGroup.POST("/users/login", h.Login)
	apiGroup.GET("/trips", middlewares.JwtAuthMiddleware(), h.GetTrip)
	apiGroup.GET("/trips/:id", middlewares.JwtAuthMiddleware(), h.GetTripById)
	apiGroup.POST("/trips", middlewares.JwtAuthMiddleware(), h.RegisterTrip)
	apiGroup.DELETE("/trips/:id", middlewares.JwtAuthMiddleware(), h.DeleteTrip)
	apiGroup.PUT("/trips/:id", middlewares.JwtAuthMiddleware(), h.UpdateTrip)
	apiGroup.POST("/transactions", middlewares.JwtAuthMiddleware(), h.RegisterTransaction)
	apiGroup.GET("/transactions/options", h.GetTransactionOptions)
	apiGroup.PUT("/transactions/:id", middlewares.JwtAuthMiddleware(), h.UpdateTransaction)
	apiGroup.DELETE("/transactions/:id", middlewares.JwtAuthMiddleware(), h.DeleteTransaction)

	engine.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return engine
}

// healthcheck
//
// @Summary			check server status
// @Description		check server status
// @Tags			default
// @Produce			json
// @Success			200	{object}	dto.BaseResponseDTO
// @Router			/v1/healthcheck [get]
func healthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"time":    time.Now().Format(time.RFC3339),
		})
	}
}

package route

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Taehoya/pocket-mate/pkg/middlewares"
	route "github.com/Taehoya/pocket-mate/pkg/routes/trip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUp(db *sql.DB, engine *gin.Engine) {
	engine.Use(gin.Recovery())
	engine.Use(middlewares.LoggingMiddleware())
	engine.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := engine.Group("api/v1")
	router.GET("healthcheck", healthCheck())
	route.NewTripRouter(db, router)
}

func healthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"time":    time.Now().Format(time.RFC3339),
		})
	}
}

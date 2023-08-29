package route

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Taehoya/pocket-mate/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func SetUp(db *sql.DB, engine *gin.Engine) {
	engine.Use(gin.Recovery())
	engine.Use(middlewares.LoggingMiddleware())

	router := engine.Group("")
	router.GET("healthcheck", healthCheck())

	NewTripRouter(db, router)

}

func healthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"time":    time.Now().Format(time.RFC3339),
		})
	}
}

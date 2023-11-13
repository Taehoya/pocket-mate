package middlewares

import (
	"net/http"

	"github.com/Taehoya/pocket-mate/internal/pkg/utils/token"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := token.IsValidToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error_message": "unathorized user",
			})
			return
		}
		ctx.Next()
	}
}

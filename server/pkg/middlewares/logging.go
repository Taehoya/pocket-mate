package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		method := ctx.Request.Method
		uri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()

		log.WithFields(log.Fields{
			"METHOD":    method,
			"URI":       uri,
			"STATUS":    statusCode,
			"LATENCY":   latency,
			"CLIENT_IP": clientIP,
		}).Info("HTTP REQUEST")

		ctx.Next()
	}

}

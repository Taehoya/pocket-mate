package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	host = os.Getenv("GO_HOST")
	port = os.Getenv("GO_PORT")
)

func init() {
	if host == "" {
		host = "localhost"
	}
	flag.StringVar(&host, "host", host, "override default hostname")
	if port == "" {
		port = "8080"
	}
	flag.StringVar(&port, "port", port, "override default port")
}

func main() {
	log.Print("starting process")

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	router.Run(host + ":" + port)
}

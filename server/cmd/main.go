package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/Taehoya/pocket-mate/docs"
	"github.com/Taehoya/pocket-mate/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

//	@title			Pocket Mate API
//	@version		1.0.0
//	@description	This is a pocket-mate server api

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	config := makeConfig()
	db, err := InitDB(config)

	if err != nil {
		log.Fatal("failed to init Database")
	}
	defer db.Close()

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.Print("starting process")

	ginHandler := gin.Default()
	route.SetUp(db, ginHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler:      ginHandler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	idleConnsClosed := make(chan struct{})
	go trapTerminationSignal(server, idleConnsClosed)

	log.Printf("services listening on %s:%s\n", config.Host, config.Port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	<-idleConnsClosed

	log.Print("end of processing")
}

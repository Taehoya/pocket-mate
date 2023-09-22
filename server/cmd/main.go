package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/Taehoya/pocket-mate/docs"
	handler "github.com/Taehoya/pocket-mate/pkg/handlers"
	tripRepository "github.com/Taehoya/pocket-mate/pkg/repositories/trip"
	tripUsecase "github.com/Taehoya/pocket-mate/pkg/usecases/trip"
	"github.com/Taehoya/pocket-mate/pkg/utils/config"
	"github.com/Taehoya/pocket-mate/pkg/utils/db"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

//	@title                  Pocket Mate API
//	@version                1.0.0
//	@description    This is a pocket-mate server api
//
// @host                localhost:8080
// @BasePath    /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	config := config.New()
	db, err := db.InitDB(config)

	if err != nil {
		log.Fatal("failed to init Database")
	}
	defer db.Close()

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.Print("starting process")

	tripRepository := tripRepository.NewTripRepository(db)
	tripUseCase := tripUsecase.NewTripUseCase(tripRepository)

	handler := handler.New(tripUseCase)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler:      handler.InitRoutes(),
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

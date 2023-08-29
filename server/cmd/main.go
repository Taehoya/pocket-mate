package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Taehoya/pocket-mate/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	db *sql.DB
)

type config struct {
	host     string
	port     string
	dbUser   string
	dbPasswd string
	dbName   string
	dbAddr   string
}

func makeConfig() *config {
	return &config{
		host:     os.Getenv("APP_URL"),
		port:     os.Getenv("PORT"),
		dbUser:   os.Getenv("DB_USERNAME"),
		dbPasswd: os.Getenv("DB_PASSWORD"),
		dbName:   os.Getenv("DB_NAME"),
		dbAddr:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
	}

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	config := makeConfig()

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.Print("starting process")

	dbConfig := mysql.Config{
		User:      config.dbUser,
		Passwd:    config.dbPasswd,
		Net:       "tcp",
		Addr:      config.dbAddr,
		DBName:    config.dbName,
		ParseTime: true,
	}

	db, err = sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	ginHandler := gin.Default()
	route.SetUp(db, ginHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", config.host, config.port),
		Handler:      ginHandler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	idleConnsClosed := make(chan struct{})
	go trapTerminationSignal(server, idleConnsClosed)

	log.Printf("services listening on %s:%s\n", config.host, config.port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	<-idleConnsClosed

	log.Print("end of processing")
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var direction string

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("failed to load .env file")
	// }

	flag.StringVar(&direction, "direction", "up", "migrate cmd: up or down")
	flag.Parse()

	dbConfig := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName:               os.Getenv("DB_NAME"),
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	databaseURL := dbConfig.FormatDSN()
	migrate, err := migrate.New("file://migration", fmt.Sprintf("mysql://%s", databaseURL))

	if err != nil {
		log.Printf("failed to init migration")
		return
	}

	if direction == "up" {
		if err := migrate.Up(); err != nil {
			log.Printf("failed to migrate up ")
		}
	}

	if direction == "down" {
		if err := migrate.Down(); err != nil {
			log.Printf("failed to migrate down")
			return
		}
	}
}

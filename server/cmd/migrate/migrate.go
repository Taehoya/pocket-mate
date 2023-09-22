package main

import (
	"flag"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var direction string

	flag.StringVar(&direction, "direction", "up", "migrate cmd: up or down")
	flag.Parse()

	migrate, err := migrate.New("file://migration", "mysql://root:root@tcp(localhost:3306)/test")

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

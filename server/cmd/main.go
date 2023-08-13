package main

import (
	"flag"
	"log"
	"os"
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
		port = "3000"
	}
	flag.StringVar(&port, "port", port, "override default port")
}

func main() {
	log.Print("starting process")

	flag.Parse()
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func trapTerminationSignal(server *http.Server, idleConnsClosed chan struct{}) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)
	<-sigChan
	log.Print("HTTP server shutdown. Draining connections...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server shutdown: %v\n", err)
	}
	close(idleConnsClosed)
}

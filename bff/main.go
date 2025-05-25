package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	endpointPopualarFonts = "/fonts/popular"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
	}
	http.HandleFunc(endpointPopualarFonts, handleGetPopularGoogleFonts)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()

	// Shutdown server on interrupt signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	// Gracefully shutdown server on interrupt
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server gracefully stopped.")

	os.Exit(0)
}

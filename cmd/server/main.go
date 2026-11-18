package main

import (
	"log"
	"net/http"
	"time"

	adapter "github.com/Sanki-funky/tenant-service/internal/adapters"
)

func main() {
	server := &http.Server{
		Addr:              ":8081",
		Handler:           adapter.NewRouter(),
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	log.Printf("Server listening on %s", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

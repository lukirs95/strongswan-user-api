package main

import (
	"log"
	"net/http"

	"github.com/lukirs95/strongswan-user-api/internal/pkg/rest"
)

func main() {
	router := rest.HandleRequests()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// // Good practice: enforce timeouts for servers you create!
		// WriteTimeout: 15 * time.Second,
		// ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

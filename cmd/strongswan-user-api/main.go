package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/lukirs95/strongswan-user-api/internal/pkg/rest"
)

func main() {
	pathToConfigfile := flag.String("s", "ipsec.secrets", "Path to ipsec.secrets")
	listeningPort := flag.String("p", "8080", "Server listening port")

	flag.Parse()

	router := rest.HandleRequests(*pathToConfigfile)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + *listeningPort,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

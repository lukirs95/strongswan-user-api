package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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

	var SESSION_KEY string = os.Getenv("SESSION_KEY")
	var ADMIN_USERNAME string = os.Getenv("ADMIN_USERNAME")
	var ADMIN_PASSWORD string = os.Getenv("ADMIN_PASSWORD")
	fmt.Println("SESSION_KEY: " + SESSION_KEY)
	fmt.Println("ADMIN_USERNAME: " + ADMIN_USERNAME)
	fmt.Println("ADMIN_PASSWORD: " + ADMIN_PASSWORD)
	fmt.Println("Serving API on Port: " + *listeningPort)
	fmt.Println("Path to secrets: " + *pathToConfigfile)
	log.Fatal(srv.ListenAndServe())
}

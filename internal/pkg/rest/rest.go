package rest

import (
	"github.com/gorilla/mux"
)

func HandleRequests() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/auth/{type}", handleAuth).Methods("GET", "POST")
	router.HandleFunc("/api/users", handleUsers).Methods("GET")
	router.HandleFunc("/api/user/{username}", handleUser).Methods("POST", "DELETE")
	return router
}

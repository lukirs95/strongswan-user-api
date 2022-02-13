package rest

import (
	"github.com/gorilla/mux"
)

func HandleRequests() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/files/md/{filename}", handleMarkdown).Methods("GET")
	router.HandleFunc("/files/md", handleMarkdownDir).Methods("GET")
	router.HandleFunc("/files/secret", handleSecretFilesDir).Methods("GET")
	router.HandleFunc("/files/secret/{filename}", handleSecretFiles).Methods("GET")
	router.HandleFunc("/api/auth/{type}", handleAuth).Methods("GET", "POST")
	router.HandleFunc("/api/users", handleUsers).Methods("GET")
	router.HandleFunc("/api/user/{username}", handleUser).Methods("POST", "DELETE")
	return router
}

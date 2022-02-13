package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lukirs95/strongswan-user-api/internal/pkg/auth"
	"github.com/lukirs95/strongswan-user-api/internal/pkg/secretfiles"
)

func handleSecretFiles(w http.ResponseWriter, r *http.Request) {
	if !auth.Authenticated(r) {
		http.Error(w, "Please login first", http.StatusForbidden)
		return
	}
	filename := mux.Vars(r)["filename"]
	html, err := secretfiles.GetFile(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(html)
}

func handleSecretFilesDir(w http.ResponseWriter, r *http.Request) {
	if !auth.Authenticated(r) {
		http.Error(w, "Please login first", http.StatusForbidden)
		return
	}
	jsonDir, err := secretfiles.GetDir()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(jsonDir))
}

package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lukirs95/strongswan-user-api/internal/pkg/auth"
)

func handleAuth(w http.ResponseWriter, r *http.Request) {
	switch mux.Vars(r)["type"] {
	case "login":
		if r.Method == "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		auth.Login(w, r)
	case "logout":
		if r.Method == "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		auth.Logout(w, r)
	case "authenticated":
		if r.Method == "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if !auth.Authenticated(r) {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}

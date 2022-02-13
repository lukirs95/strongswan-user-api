package auth

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var SESSION_KEY string = os.Getenv("SESSION_KEY")
var ADMIN_USERNAME string = os.Getenv("ADMIN_USERNAME")
var ADMIN_PASSWORD string = os.Getenv("ADMIN_PASSWORD")

var store = sessions.NewCookieStore([]byte(SESSION_KEY))

type admin_user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "vpn-admin")

	// Authentication goes here
	user := new(admin_user)
	json.NewDecoder(r.Body).Decode(user)

	if user.Username == ADMIN_USERNAME && user.Password == ADMIN_PASSWORD {
		// Set user as authenticated
		session.Values["authenticated"] = true
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		session.Values["authenticated"] = false
	}
	session.Save(r, w)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "vpn-admin")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func Authenticated(r *http.Request) bool {
	session, _ := store.Get(r, "vpn-admin")
	return session.Values["authenticated"].(bool)
}

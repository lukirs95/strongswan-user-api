package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lukirs95/strongswan-user-api/internal/pkg/auth"
	"github.com/lukirs95/strongswan-user-api/internal/pkg/strongswanservice"
	"github.com/lukirs95/strongswan-user-api/pkg/configparser"
	"github.com/lukirs95/strongswan-user-api/pkg/strongswanuser"
)

type vpn_user struct {
	Username string
	Password string `json:"password"`
}

func readConfig(path string) (*strongswanuser.List, error) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return configparser.Deserialize(file)
}

func saveConfig(path string, list *strongswanuser.List) error {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	err = configparser.Serialize(file, list)
	file.Close()
	return err
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	if !auth.Authenticated(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	list, err := readConfig("ipsec.secrets")
	if err != nil {
		http.Error(w, "could not open config file", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(list.Json()))
	if err != nil {
		fmt.Println(err)
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	if !auth.Authenticated(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	vpn_user := &vpn_user{Username: mux.Vars(r)["username"]}
	list, err := readConfig("ipsec.secrets")
	if err != nil {
		http.Error(w, "could not open config file", http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case "POST":
		json.NewDecoder(r.Body).Decode(vpn_user)
		err = list.Append(vpn_user.Username, vpn_user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = saveConfig("ipsec.secrets", list)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "DELETE":
		err = list.Remove(vpn_user.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = saveConfig("ipsec.secrets", list)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	err = strongswanservice.Restart()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

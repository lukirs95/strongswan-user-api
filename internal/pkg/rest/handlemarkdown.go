package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lukirs95/strongswan-user-api/internal/pkg/mdtohtml"
)

func handleMarkdown(w http.ResponseWriter, r *http.Request) {
	filename := mux.Vars(r)["filename"]
	html, err := mdtohtml.GetFile(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(html)
}

func handleMarkdownDir(w http.ResponseWriter, r *http.Request) {
	jsonDir, err := mdtohtml.GetDir()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(jsonDir))
}

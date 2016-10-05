package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	// Gorilla mux router
	r := mux.NewRouter().StrictSlash(true)

	// Handlers
	r.HandleFunc("/", mainHandler)
	r.HandleFunc("/healtz", healtzHander)
	r.HandleFunc("/version", versionHandler)

	return r
}

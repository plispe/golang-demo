package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter - Constructor
func NewRouter() http.Handler {
	// Gorilla mux router
	r := mux.NewRouter().StrictSlash(true)

	// Handlers
	r.PathPrefix("/socket.io/").Handler(NewSocketIOHandler())
	r.HandleFunc("/healtz", healtzHander)
	r.HandleFunc("/version", versionHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	return r
}

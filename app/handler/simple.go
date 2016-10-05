package handler

import (
	"fmt"
	"net/http"
)

const (
	version = "1.0.0-dev"
)

func MainHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func HealtzHander(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Yean, I'm ready.")
}

func VersionHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, version)
}

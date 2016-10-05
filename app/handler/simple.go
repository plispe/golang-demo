package handler

import (
	"fmt"
	"net/http"
)

const (
	version = "1.0.0-dev"
)

func mainHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func healtzHander(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Yean, I'm ready.")
}

func versionHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, version)
}

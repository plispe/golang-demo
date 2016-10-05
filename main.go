package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

const (
	version = "1.0.0-dev"
)

func main() {
	// Gorilla mux router
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello world!")
	})

	r.HandleFunc("/healtz", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Yean, I'm ready.")
	})

	r.HandleFunc("/version", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, version)
	})

	// Middleware dispatcher
	n := negroni.New()
	n.UseHandler(r)

	// Load port from envirnment
	port := os.Getenv("PORT")

	// Assing default port if not given
	if port == "" {
		port = "8888"
	}

	// Construct server address
	address := fmt.Sprintf(":%s", port)

	// Log on stdout
	log.Println(fmt.Sprintf(`Server is listening on address "%s"`, address))

	// Start server
	http.ListenAndServe(address, n)
}

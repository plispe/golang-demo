package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/plispe/golang-demo/app/handler"
	"github.com/urfave/negroni"
)

func main() {
	// Gorilla mux router
	r := mux.NewRouter().StrictSlash(true)

	// Handlers
	r.HandleFunc("/", handler.MainHandler)
	r.HandleFunc("/healtz", handler.HealtzHander)
	r.HandleFunc("/version", handler.VersionHandler)

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

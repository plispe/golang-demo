package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/plispe/golang-demo/app/handler"
	"github.com/plispe/golang-demo/app/middleware"
)

func main() {
	// Middleware dispatcher
	d := middleware.NewDispatcher()
	// Use gorilla mux as negroni handler
	d.UseHandler(handler.NewRouter())

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
	http.ListenAndServe(address, d)
}

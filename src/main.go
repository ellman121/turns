package main

import (
	"log"
	"net/http"
	"routes"
)

var port = ":5500"

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", routes.Home)

	http.HandleFunc("/", routes.Home)

	log.Printf("Listening on port (%s)", port)
	http.ListenAndServe(port, nil)
}

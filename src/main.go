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
	http.HandleFunc("/create", create)
	http.HandleFunc("/join", join)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Listening on port (%s)", port)
	http.ListenAndServe(port, nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	log.Println("[create] [requestIP " + r.RemoteAddr + "]")
}

func join(w http.ResponseWriter, r *http.Request) {
	log.Println("[join] [requestIP " + r.RemoteAddr + "]")
}

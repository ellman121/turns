package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"routes"
)

var port = ":5500"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", routes.Home)
	mux.HandleFunc("/create", create)
	mux.HandleFunc("/join", join)
	mux.HandleFunc("/newSession", newID)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Listening on port (%s)", port)
	http.ListenAndServe(port, mux)
}

func create(w http.ResponseWriter, r *http.Request) {
	log.Println("[create] [requestIP " + r.RemoteAddr + "]")
}

func join(w http.ResponseWriter, r *http.Request) {
	log.Println("[join] [requestIP " + r.RemoteAddr + "]")
}

func newID(w http.ResponseWriter, r *http.Request) {
	log.Println("[newID] [requestIP " + r.RemoteAddr + "]")

	ID := rand.Uint32() % 99999

	fmt.Fprintf(w, "{gameID: %d}", ID)
}

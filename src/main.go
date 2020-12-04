package main

import (
	"log"
	"net/http"

	"models"
	"routes"
)

var port = ":5500"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", routes.Home)
	mux.HandleFunc("/create", create)
	mux.HandleFunc("/join", routes.Join)
	mux.HandleFunc("/newSession", routes.NewSession)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Listening on port (%s)", port)

	go func() {
		s, err := models.NewSession()
		if err != nil {
			log.Println("[main] error creating default session")
			return
		}
		log.Println("[main] Session created with id " + s.ID)
	}()

	http.ListenAndServe(port, mux)
}

func create(w http.ResponseWriter, r *http.Request) {
	log.Println("[create] [requestIP " + r.RemoteAddr + "]")
}

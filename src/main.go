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

	// REST API
	mux.HandleFunc("/newGame", routes.NewGame)

	// SOCKET API
	mux.HandleFunc("/ws", routes.SocketHandler)

	// HTTP ROUTES
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", routes.Home)

	// Create a defualt game for testing purposes
	go func() {
		s, err := models.NewGame()
		if err != nil {
			log.Println("[main] error creating default game")
			return
		}
		log.Println("[main] Game created with id " + s.ID)
	}()

	log.Printf("Listening on port (%s)", port)
	http.ListenAndServe(port, mux)
}

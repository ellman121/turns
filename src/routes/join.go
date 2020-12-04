package routes

import (
	"log"
	"models"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Join - Join a session
func Join(w http.ResponseWriter, r *http.Request) {
	log.Println("[join] [requestIP " + r.RemoteAddr + "]")

	err := r.ParseForm()
	if err != nil {
		log.Println("[join] error parsing form values from URL")
		return
	}

	gameID := r.Form.Get("gameID")
	if gameID == "" {
		return
	}

	s, err := models.GetSession(gameID)
	if err != nil {
		log.Println("[join] unable to find game with ID " + gameID)
		return
	}

	err = rnd.JSON(w, http.StatusOK, map[string]interface{}{})
	log.Printf("%v", s)
}

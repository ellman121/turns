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

// SocketHandler - SocketHandler a session
func SocketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[join] [requestIP " + r.RemoteAddr + "]")

	err := r.ParseForm()
	if err != nil {
		log.Println("[join] error parsing form values from URL")
		rnd.JSON(w, http.StatusBadRequest, map[string]interface{}{})
		return
	}

	gameID := r.Form.Get("gameID")
	if gameID == "" {
		log.Println("[join] No gameID passed from client")
		rnd.JSON(w, http.StatusBadRequest, map[string]interface{}{})
		return
	}

	s, err := models.GetSession(gameID)
	if err != nil {
		log.Println("[join] unable to find game with ID " + gameID)
		rnd.JSON(w, http.StatusNotFound, map[string]interface{}{})
		return
	}

	log.Println("[join] Upgrading connection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[join] error upgrading connection %v", err)
		return
	}

	log.Println(s)
	log.Println(conn.RemoteAddr())
}

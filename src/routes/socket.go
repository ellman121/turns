package routes

import (
	"encoding/json"
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
	log.Println("[SocketHandler] [requestIP " + r.RemoteAddr + "]")

	err := r.ParseForm()
	if err != nil {
		log.Println("[SocketHandler] error parsing form values from URL")
		rnd.JSON(w, http.StatusBadRequest, map[string]interface{}{})
		return
	}

	gameID := r.Form.Get("gameID")
	if gameID == "" {
		log.Println("[SocketHandler] No gameID passed from client")
		rnd.JSON(w, http.StatusBadRequest, map[string]interface{}{})
		return
	}

	s, err := models.GetSession(gameID)
	if err != nil {
		log.Println("[SocketHandler] unable to find game with ID " + gameID)
		rnd.JSON(w, http.StatusNotFound, map[string]interface{}{})
		return
	}

	log.Println("[SocketHandler] Upgrading connection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[SocketHandler] error upgrading connection %v", err)
		return
	}

	go manageSocketConnection(conn, s)
}

func manageSocketConnection(conn *websocket.Conn, session *models.Session) {
	defer conn.Close()

	// On connection, we instantly send the current state of the game
	gameState, err := json.Marshal(session)
	if err != nil {
		log.Printf("[manageSocket] [%v] Failed to marshal game state", session.ID)
	}

	w, err := conn.NextWriter(websocket.TextMessage)
	if err != nil {
		log.Printf("[manageSocket] [%v] Failed to send initial game state to client\n", session.ID)
	}
	w.Write(gameState)
	w.Close()

	for {
		messageType, _, err := conn.NextReader()
		if err != nil {
			log.Println("[manageSocket] Error creating nextReader")
			return
		}

		w, err := conn.NextWriter(messageType)
		if err != nil {
			log.Println("[manageSocket] Error creating nextWriter")
		}

		w.Write([]byte("Hello World"))
		w.Close()
	}
}

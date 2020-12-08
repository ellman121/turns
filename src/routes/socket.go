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

	s, err := models.GetGame(gameID)
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

func manageSocketConnection(conn *websocket.Conn, session *models.Game) {
	defer conn.Close()

	// Generate a secure player ID and save it in the cache

	// On connection, we instantly send the current state of the game
	s, err := json.Marshal(session)
	if err != nil {
		log.Printf("[manageSocket] [%v] Failed to marshal game state", session.ID)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, s); err != nil {
		log.Printf("[manageSocket] [%v] Failed to send initial game state to client\n", session.ID)
		return
	}

	// for {
	// 	err := conn.ReadJSON()
	// 	if err != nil {
	// 		log.Println("[manageSocket] Error reading JSON from socket")
	// 		return
	// 	}
	// 	r.Read()

	// 	w, err := conn.NextWriter(messageType)
	// 	if err != nil {
	// 		log.Println("[manageSocket] Error creating nextWriter")
	// 	}

	// 	w.Write([]byte("Hello World"))
	// 	w.Close()
	// }
}

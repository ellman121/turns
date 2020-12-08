package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"models"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// SocketHandler - Upgrade and connect to socket connections
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

	g, err := models.GetGame(gameID)
	if err != nil {
		log.Printf("[SocketHandler] %v", err)
		rnd.JSON(w, http.StatusNotFound, map[string]interface{}{})
		return
	}

	log.Println("[SocketHandler] Upgrading connection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[SocketHandler] error upgrading connection %v", err)
		return
	}

	go manageSocketConnection(conn, g)
}

func manageSocketConnection(conn *websocket.Conn, game *models.Game) {
	defer conn.Close()

	// Generate a player ID and save it to the game.
	err := game.AddPlayer(fmt.Sprintf("%d", rand.Intn(900000)+100000))
	if err != nil {
		log.Printf("[manageSocket] [%v] %v", game.ID, err)
		conn.WriteControl(websocket.CloseMessage, []byte{}, time.Now())
		conn.ReadMessage()
		conn.Close()
		return
	}

	// On connection, we instantly send the current state of the game
	g, err := json.Marshal(game)
	if err != nil {
		log.Printf("[manageSocket] [%v] Failed to marshal game state", game.ID)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, g); err != nil {
		log.Printf("[manageSocket] [%v] Failed to send initial game state to client\n", game.ID)
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

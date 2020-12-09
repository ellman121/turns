package routes

import (
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

	log.Println("[SocketHandler] Upgrading connection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[SocketHandler] error upgrading connection %v", err)
		return
	}

	go manageSocketConnection(conn, gameID)
}

func manageSocketConnection(conn *websocket.Conn, gameID string) {
	defer func() {
		conn.WriteControl(websocket.CloseMessage, []byte{}, time.Now().Add(1))
		conn.Close()
	}()

	game, err := models.GetGame(gameID)
	if err != nil {
		log.Printf("[manageSocket %v] Failed to get game by ID", gameID)
		return
	}

	// Generate a player ID and save it to the game.
	playerID := fmt.Sprintf("%d", rand.Intn(900000)+100000)
	err = game.AddPlayer(playerID)
	if err != nil {
		log.Printf("[manageSocket %v] %v", gameID, err)
		return
	}

	// On connection, we instantly send the current state of the game
	if err = conn.WriteJSON(game.Sanatized(playerID)); err != nil {
		log.Printf("[manageSocket %v] Failed to send initial game state to client\n", gameID)
		return
	}

	for {
		var t models.Transform
		err := conn.ReadJSON(&t)
		if err != nil {
			if err.Error() == "websocket: close 1000 (normal)" {
				return
			}

			log.Println("[manageSocket] Error reading JSON from socket")
			return
		}

		// Get the latest version of the game
		game, err = models.ProcessTransform(t)
		if err != nil {
			log.Printf("[manageSocket %v] %v", gameID, err)
			return
		}

		// Sanatize the output
		if game.Players[0].ID == playerID {
			game.Players[1].ID = ""
		} else {
			game.Players[0].ID = ""
		}

		if err = conn.WriteJSON(game.Sanatized(playerID)); err != nil {
			log.Printf("[manageSocket %v] Failed to send updated game state to client\n", gameID)
		}
	}
}

package routes

import (
	"log"
	"math/rand"
	"net/http"

	"models"
)

// NewSession - Create a new session and return the details via JSON
func NewSession(w http.ResponseWriter, r *http.Request) {
	log.Println("[newID] [requestIP " + r.RemoteAddr + "]")

	ID := rand.Uint32() % 99999

	rnd.JSON(w, http.StatusOK, models.Session{
		ID: ID,
	})
}

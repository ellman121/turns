package routes

import (
	"log"
	"net/http"

	"models"
)

// NewGame - Create a new game and return the details via JSON
func NewGame(w http.ResponseWriter, r *http.Request) {
	log.Println("[newID] [requestIP " + r.RemoteAddr + "]")

	s, err := models.NewGame()
	if err != nil {
		rnd.HTML(w, http.StatusInternalServerError, "5XX", nil)
		return
	}

	rnd.JSON(w, http.StatusOK, *s)
}

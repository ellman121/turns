package routes

import (
	"log"
	"net/http"
)

// Home - Render the home page
func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving home page " + r.RemoteAddr)

	err := rnd.HTML(w, http.StatusOK, "home", nil)
	if err != nil {
		log.Fatal(err)
	}
}

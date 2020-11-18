package routes

import (
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	rnd = renderer.New(renderer.Options{
		ParseGlobPattern: "routes/templates/*.html",
	})
}

// Home - Render the home page
func Home(w http.ResponseWriter, r *http.Request) {
	err := rnd.HTML(w, http.StatusOK, "home", nil)

	if err != nil {
		log.Fatal(err)
	}
}

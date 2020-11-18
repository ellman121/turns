package routes

import "github.com/thedevsaddam/renderer"

var rnd *renderer.Render

func init() {
	rnd = renderer.New(renderer.Options{
		ParseGlobPattern: "html/*.html",
	})
}

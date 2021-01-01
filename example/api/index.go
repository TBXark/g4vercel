package handler

import (
	"github.com/tbxark/gee4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := gee.Default()
	server.Get("/", func(context *Context) {
		c.String(200, "It's work\n")
	})
	server.Handler(w, r)
}

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/webdevfuel/gopizza/template"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		component := template.Hello("world")
		component.Render(r.Context(), w)
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Hx-Trigger", `{"add-toast": {"message": "Successfully triggered from backend.", "type": "info"}}`)
		component := template.Action("sent")
		component.Render(r.Context(), w)
	})
	http.ListenAndServe("localhost:3000", r)
}

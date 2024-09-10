package common

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iypetrov/gopizza/templates"
	"net/http"
)

func NewRouter(ctx context.Context) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/", func(r chi.Router) {
		r.Handle("/web/*", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			component := templates.Hello("world")
			component.Render(r.Context(), w)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Hx-Trigger", `{"add-toast": {"message": "Successfully triggered from backend.", "type": "info"}}`)
			component := templates.Action("sent")
			component.Render(r.Context(), w)
		})
	})

	r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte{})
		if err != nil {
			return
		}
	})

	return r
}

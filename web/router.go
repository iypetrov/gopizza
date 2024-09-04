package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/iypetrov/gopizza/internal/middleware"
	"github.com/iypetrov/gopizza/internal/utils"
	"net/http"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Get("/404", NotFoundHandler)
	r.Get("/home", HomeHandler)
	r.With(middleware.UUIDOrNotFound).Get("/pizza/{id}", PizzaHandler)
	r.NotFound(utils.RedirectHomeView)

	return r
}

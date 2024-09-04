package web

import (
	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/config"
	"github.com/iypetrov/gopizza/internal/middleware"
	"github.com/iypetrov/gopizza/web/templates/views"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.NotFound(config.Get().GetAPIPrefix())).ServeHTTP(w, r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.Home(config.Get().GetAPIPrefix())).ServeHTTP(w, r)
}

func PizzaHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := r.Context().Value(middleware.UUIDKey).(uuid.UUID)
	templ.Handler(views.Pizza(config.Get().GetAPIPrefix(), id)).ServeHTTP(w, r)
}

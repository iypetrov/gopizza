package pizzas

import (
	"github.com/go-chi/chi/v5"
	"github.com/iypetrov/gopizza/internal/utils"
	"net/http"
)

func Router(handler *PizzaHandler) http.Handler {
	r := chi.NewRouter()
	r.Post("/", utils.MakeHandler(handler.createPizza))
	return r
}

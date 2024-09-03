package pizzas

import (
	"github.com/go-chi/chi/v5"
	"github.com/iypetrov/gopizza/internal/middleware"
	"github.com/iypetrov/gopizza/internal/utils"
	"net/http"
)

func Router(handler *PizzaHandler) http.Handler {
	r := chi.NewRouter()
	r.With(middleware.BodyFormat).Post("/", utils.MakeHandler(handler.createPizza))
	r.With(middleware.UUIDFormat).Get("/{id}", utils.MakeHandler(handler.getPizzaByID))
	r.Get("/", utils.MakeHandler(handler.getAllPizzas))
	r.With(middleware.UUIDFormat).With(middleware.BodyFormat).Put("/{id}", utils.MakeHandler(handler.updatePizza))
	r.With(middleware.UUIDFormat).Delete("/{id}", utils.MakeHandler(handler.deletePizzaByID))
	return r
}

package router

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/iypetrov/gopizza/configs"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/handlers"
	"github.com/iypetrov/gopizza/internal/services"
)

func NewRouter(ctx context.Context, db *database.Queries) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// services
	pizzaSrv := services.NewPizza(ctx, db)

	// handlers
	authHnd := handlers.NewAuth()
	pizzaHnd := handlers.NewPizza(pizzaSrv)

	r.Route("/", func(r chi.Router) {
		// common
		r.Handle("/web/*", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
		r.Get("/404", Make(handlers.NotFoundView))

		// client
		r.Get("/home", Make(handlers.HomeView))
		r.Get("/login", Make(handlers.LoginView))
		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			handlers.RedirectHomePage(w)
		})

		// admin
		r.Route(configs.Get().GetAdminPrefix(), func(r chi.Router) {
			r.Get("/home", Make(handlers.AdminHomeView))
			r.NotFound(func(w http.ResponseWriter, r *http.Request) {
				handlers.RedirectAdminHomePage(w)
			})
		})

		// api
		r.Route(configs.Get().GetAPIPrefix(), func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Post("/login", Make(authHnd.Login))
				r.Route("/pizzas", func(r chi.Router) {
					r.Post("/", Make(pizzaHnd.CreatePizza))
					r.Get("/{id}", Make(pizzaHnd.GetPizzaByID))
					r.Get("/admin/overview", Make(pizzaHnd.GetAllPizzasAdminOverview))
					r.Put("/{id}", Make(pizzaHnd.UpdatePizza))
					r.Delete("/{id}", Make(pizzaHnd.DeletePizzaByID))
				})
			})
		})

		r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte{})
			if err != nil {
				return
			}
		})
	})

	return r
}

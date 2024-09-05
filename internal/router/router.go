package router

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/iypetrov/gopizza/internal/config"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/handler"
	mid "github.com/iypetrov/gopizza/internal/middleware"
	"github.com/iypetrov/gopizza/internal/repository"
	"github.com/iypetrov/gopizza/internal/service"
	"github.com/iypetrov/gopizza/internal/util"
	"github.com/iypetrov/gopizza/web"
	"net/http"
)

func New(ctx context.Context, db *database.Queries) *chi.Mux {
	// repositories
	pizzaRep := repository.NewPizza(db)

	// services
	pizzaSrv := service.NewPizza(ctx, pizzaRep)

	// handlers
	pizzaHnd := handler.NewPizza(pizzaSrv)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	if config.Get().App.Environment == config.DevEnv {
		r.Use(middleware.Logger)
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: config.Get().App.Environment != config.DevEnv,
		MaxAge:           300,
	}))

	r.Mount("/", web.Router())
	r.Route(config.Get().GetAPIPrefix(), func(r chi.Router) {
		r.Use(GetCtxVersion(config.Get().App.Version))

		// public routes
		r.Group(func(r chi.Router) {
			r.Route("/pizzas", func(r chi.Router) {
				r.
					With(mid.BodyFormat).
					Post("/", util.Make(pizzaHnd.CreatePizza))
				r.
					With(mid.UUIDFormat).
					Get("/{id}", util.Make(pizzaHnd.GetPizzaByID))
				r.
					Get("/", util.Make(pizzaHnd.GetAllPizzas))
				r.
					With(mid.UUIDFormat).
					With(mid.BodyFormat).
					Put("/{id}", util.Make(pizzaHnd.UpdatePizza))
				r.
					With(mid.UUIDFormat).
					Delete("/{id}", util.Make(pizzaHnd.DeletePizzaByID))
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

func GetCtxVersion(version string) func(next http.Handler) http.Handler {
	versionKey := "API_VERSION"
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), versionKey, version))
			next.ServeHTTP(w, r)
		})
	}
}

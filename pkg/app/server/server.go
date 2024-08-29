package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/iypetrov/gopizza/pkg/app"
	"github.com/iypetrov/gopizza/pkg/app/config"
	"github.com/iypetrov/gopizza/pkg/pizzas"
	"net/http"
	"time"
)

func New() *http.Server {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", app.Cfg.App.Port),
		Handler:      registerRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func registerRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	if app.Cfg.App.Environment == config.DevEnv {
		r.Use(middleware.Logger)
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: app.Cfg.App.Environment != config.DevEnv,
		MaxAge:           300,
	}))

	r.Route(fmt.Sprintf("/api/v%s", app.Cfg.App.Version), func(r chi.Router) {
		r.Use(apiVersionCtx(app.Cfg.App.Version))
		// Public Routes
		r.Group(func(r chi.Router) {
			r.Mount("/pizzas", pizzas.Router())
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


func apiVersionCtx(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "api.version", version))
			next.ServeHTTP(w, r)
		})
	}
}

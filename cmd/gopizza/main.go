package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/iypetrov/gopizza/internal/config"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/pizzas"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
	cfg    *config.Config
	log    *config.Logger
	db     *database.Queries

	pizzasHnd *pizzas.PizzaHandler
)

func init() {
	ctx, cancel = context.WithCancel(context.Background())
}

func main() {
	cfg = config.New()
	log = config.NewLogger()
	conn, err := config.CreateDatabaseConnection(cfg)
	if err != nil {
		log.Error("cannot connect to database %s", err.Error())
	}
	db = config.NewDatabase(conn)
	if err := config.RunSchemaMigration(conn); err != nil {
		log.Error("cannot run schema migration %s", err.Error())
	}

	// repositories
	pizzasRep := pizzas.NewRepository(db)

	// services
	pizzasSrv := pizzas.NewService(ctx, log, pizzasRep)

	// handlers
	pizzasHnd = pizzas.NewHandler(pizzasSrv)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.App.Port),
		Handler:      registerRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Info("server started on %s\n", cfg.App.Port)
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	<-setupGracefulShutdown(cancel)
}

func registerRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	if cfg.App.Environment == config.DevEnv {
		r.Use(middleware.Logger)
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: cfg.App.Environment != config.DevEnv,
		MaxAge:           300,
	}))

	r.Route(fmt.Sprintf("/api/v%s", cfg.App.Version), func(r chi.Router) {
		r.Use(apiVersionCtx(cfg.App.Version))
		// Public Routes
		r.Group(func(r chi.Router) {
			r.Mount("/pizzas", pizzas.Router(pizzasHnd))
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

func setupGracefulShutdown(cancel context.CancelFunc) (shutdownCompleteChan chan struct{}) {
	shutdownCompleteChan = make(chan struct{})
	isFirstShutdownSignal := true

	shutdownFunc := func() {
		if !isFirstShutdownSignal {
			log.Info("caught another exit signal, now hard dying")
			os.Exit(1)
		}

		isFirstShutdownSignal = false
		log.Info("starting graceful shutdown")

		cancel()

		close(shutdownCompleteChan)
	}

	go func(shutdownFunc func()) {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		for {
			log.Info("caught exit signal", "signal", <-sigint)
			go shutdownFunc()
		}
	}(shutdownFunc)

	return shutdownCompleteChan
}

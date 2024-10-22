package router

import (
	"context"
	"database/sql"
	"net/http"

	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/iypetrov/gopizza/configs"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/handlers"
	"github.com/iypetrov/gopizza/internal/middlewares"
	"github.com/iypetrov/gopizza/internal/services"
)

func NewRouter(ctx context.Context, db *sql.DB, queries *database.Queries, s3Client *s3.Client, cognitoClient *cip.Client) *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Recoverer)
	// mux.Use(middleware.Logger)

	// services
	imageSrv := services.NewImage(s3Client)
	authSrv := services.NewAuth(db, queries, cognitoClient)
	pizzaSrv := services.NewPizza(db, queries)

	// handlers
	imageHnd := handlers.NewImage(imageSrv)
	authHnd := handlers.NewAuth(authSrv)
	pizzaHnd := handlers.NewPizza(pizzaSrv, imageSrv)

	mux.Route("/", func(mux chi.Router) {
		// common
		mux.Handle("/web/*", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
		mux.Get("/404", Make(handlers.NotFoundView))
		mux.With(middlewares.UUIDFormat).Get("/image/{id}", imageHnd.GetImage)
		mux.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte{})
			if err != nil {
				return
			}
		})

		// client
		mux.Get("/home", Make(handlers.HomeView))
		mux.Get("/register", Make(handlers.RegisterView))
		mux.Get("/verification-code", Make(handlers.RegisterVerificationView))
		mux.Get("/login", Make(handlers.LoginView))

		// admin
		mux.Route(configs.Get().GetAdminPrefix(), func(mux chi.Router) {
			mux.Get("/home", Make(handlers.AdminHomeView))
		})

		// client api
		mux.Route(configs.Get().GetClientAPIPrefix(), func(mux chi.Router) {
			mux.Group(func(r chi.Router) {
				r.Post("/register", Make(authHnd.Register))
				r.Post("/verification-code", Make(authHnd.VerifyRegistrationCode))
				r.Post("/login", Make(authHnd.Login))
			})
		})

		// admin api
		mux.Route(configs.Get().GetAdminAPIPrefix(), func(mux chi.Router) {
			mux.Group(func(r chi.Router) {
				r.Route("/pizzas", func(r chi.Router) {
					r.Post("/", Make(pizzaHnd.AdminCreatePizza))
					r.Get("/", Make(pizzaHnd.AdminGetAllPizzas))
					r.With(middlewares.UUIDFormat).Delete("/{id}", Make(pizzaHnd.AdminDeletePizzaByID))
				})
			})
		})
	})

	return mux
}

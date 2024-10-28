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

func NewRouter(ctx context.Context, public http.Handler, db *sql.DB, queries *database.Queries, s3Client *s3.Client, cognitoClient *cip.Client) *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Recoverer)
	// mux.Use(middleware.Logger)

	// services
	imageSrv := services.NewImage(s3Client)
	authSrv := services.NewAuth(db, queries, cognitoClient)
	pizzaSrv := services.NewPizza(db, queries)
	cartSrv := services.NewCart(db, queries)

	// handlers
	imageHnd := handlers.NewImage(imageSrv)
	authHnd := handlers.NewAuth(authSrv)
	pizzaHnd := handlers.NewPizza(pizzaSrv, imageSrv)
	cartHnd := handlers.NewCart(cartSrv)

	mux.Route("/", func(mux chi.Router) {
		// common
		mux.Handle("/public/*", public)
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
		mux.Get("/register", Make(handlers.RegisterView))
		mux.Get("/verification-code", Make(handlers.RegisterVerificationView))
		mux.Get("/login", Make(handlers.LoginView))
		mux.Get("/home", Make(handlers.HomeView))
		mux.Get("/checkout", Make(handlers.CheckoutView))
		mux.With(middlewares.UUIDFormat).Get("/pizzas/{id}", Make(handlers.PizzaDetailsView))

		// admin
		mux.Route(configs.Get().AdminPrefix(), func(mux chi.Router) {
			mux.Get("/home", Make(handlers.AdminHomeView))
		})

		// public api
		mux.Route(configs.Get().PublicAPIPrefix(), func(mux chi.Router) {
			mux.Group(func(r chi.Router) {
				r.Post("/register", Make(authHnd.Register))
				r.Post("/verification-code", Make(authHnd.VerifyRegistrationCode))
				r.Post("/login", Make(authHnd.Login))
				r.Post("/logout", Make(authHnd.Logout))
			})
		})

		// client api
		mux.With(middlewares.AuthClient).Route(configs.Get().ClientAPIPrefix(), func(mux chi.Router) {
			mux.Route("/pizzas", func(mux chi.Router) {
				mux.Get("/", Make(pizzaHnd.GetAllPizzas))
				mux.With(middlewares.UUIDFormat).Get("/{id}", Make(pizzaHnd.GetPizzaByID))
			})
			mux.With(middlewares.AuthClient).Route("/carts", func(mux chi.Router) {
				mux.With(middlewares.UUIDFormat).Post("/pizzas/{id}", Make(cartHnd.AddPizzaToCart))
				mux.Get("/", Make(cartHnd.GetCartByUserID))
				mux.Delete("/", Make(cartHnd.EmptyCartByUserID))
				mux.With(middlewares.UUIDFormat).Delete("/{id}", Make(cartHnd.RemoveItemFromCart))
			})
		})

		// admin api
		mux.With(middlewares.AuthAdmin).Route(configs.Get().AdminAPIPrefix(), func(mux chi.Router) {
			mux.Group(func(mux chi.Router) {
				mux.Route("/pizzas", func(mux chi.Router) {
					mux.Post("/", Make(pizzaHnd.AdminCreatePizza))
					mux.Get("/", Make(pizzaHnd.AdminGetAllPizzas))
					mux.With(middlewares.UUIDFormat).Delete("/{id}", Make(pizzaHnd.AdminDeletePizzaByID))
				})
			})
		})
	})

	return mux
}

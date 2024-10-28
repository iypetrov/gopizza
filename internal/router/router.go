package router

import (
	"context"
	"database/sql"
	"net/http"

	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-chi/chi/v5"

	"github.com/iypetrov/gopizza/configs"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/handlers"
	"github.com/iypetrov/gopizza/internal/middlewares"
	"github.com/iypetrov/gopizza/internal/services"
)

func NewRouter(ctx context.Context, public http.Handler, db *sql.DB, queries *database.Queries, s3Client *s3.Client, cognitoClient *cip.Client) *chi.Mux {
	mux := chi.NewRouter()

	// services
	imageSrv := services.NewImage(s3Client)
	authSrv := services.NewAuth(db, queries, cognitoClient)
	pizzaSrv := services.NewPizza(db, queries)
	cartSrv := services.NewCart(db, queries)
	orderSrv := services.NewOrder(db, queries)
	paymentSrv := services.NewPayment(orderSrv)

	// handlers
	imageHnd := handlers.NewImage(imageSrv)
	authHnd := handlers.NewAuth(authSrv)
	pizzaHnd := handlers.NewPizza(pizzaSrv, imageSrv)
	cartHnd := handlers.NewCart(cartSrv)
	orderHnd := handlers.NewOrder(orderSrv)
	paymentHnd := handlers.NewPayment(paymentSrv)

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
		mux.Get("/checkout/tracking", Make(handlers.TrackingView))
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
			mux.Route("/carts", func(mux chi.Router) {
				mux.With(middlewares.UUIDFormat).Post("/pizzas/{id}", Make(cartHnd.AddPizzaToCart))
				mux.Get("/", Make(cartHnd.GetCartByUserID))
				mux.Delete("/", Make(cartHnd.EmptyCartByUserID))
				mux.With(middlewares.UUIDFormat).Delete("/{id}", Make(cartHnd.RemoveItemFromCart))
			})
			mux.Route("/orders", func(mux chi.Router) {
				mux.Post("/", Make(orderHnd.CreateOrder))
				mux.Get("/", Make(orderHnd.GetOrderByIntentID))
			})
			mux.Route("/payments", func(mux chi.Router) {
				mux.Get("/config", Make(paymentHnd.GetPublishableKey))
				mux.Post("/metadata", Make(paymentHnd.GetPaymentMetadata))
				mux.Post("/webhook", Make(paymentHnd.HandleWebhookEvent))
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

package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/middlewares"
	"github.com/iypetrov/gopizza/internal/toasts"
	"github.com/iypetrov/gopizza/templates/views"
)

func NotFoundView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.NotFound())
}

func RegisterView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.Register())
}

func RegisterVerificationView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.RegisterVerification(*r))
}

func LoginView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.Login())
}

func HomeView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.Home())
}

func AdminHomeView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.AdminHome())
}

func CheckoutView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.Checkout())
}

func TrackingView(w http.ResponseWriter, r *http.Request) error {
	intentID := r.URL.Query().Get("payment_intent")
	return Render(w, r, views.Tracking(intentID))
}

func PizzaDetailsView(w http.ResponseWriter, r *http.Request) error {
	id, ok := r.Context().Value(middlewares.UUIDKey).(uuid.UUID)
	if !ok {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidUUID))
		return toasts.ErrNotValidUUID
	}
	return Render(w, r, views.PizzaDetails(id))
}


func SaladDetailsView(w http.ResponseWriter, r *http.Request) error {
	id, ok := r.Context().Value(middlewares.UUIDKey).(uuid.UUID)
	if !ok {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidUUID))
		return toasts.ErrNotValidUUID
	}
	return Render(w, r, views.SaladDetails(id))
}
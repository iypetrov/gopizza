package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/dtos"
	"github.com/iypetrov/gopizza/internal/middlewares"
	"github.com/iypetrov/gopizza/internal/services"
	"github.com/iypetrov/gopizza/internal/toasts"
)

type Payment struct {
	srv services.Payment
}

func NewPayment(srv services.Payment) Payment {
	return Payment{
		srv: srv,
	}
}

func (hnd *Payment) GetPublishableKey(w http.ResponseWriter, r *http.Request) error {
	resp := dtos.PaymentPublishableKeyResponse{
		PublishableKey: hnd.srv.GetPublishableKey(),
	}
	return WriteJson(w, http.StatusOK, resp)
}

func (hnd *Payment) CreateIntent(w http.ResponseWriter, r *http.Request) error {
	cookie, ok := r.Context().Value(middlewares.CookieName).(dtos.UserCookie)
	if !ok {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidCookie))
		return toasts.ErrNotValidCookie
	}

	userID, err := uuid.Parse(cookie.ID)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidUUID))
		return toasts.ErrorInternalServerError(toasts.ErrNotValidUUID)
	}

	var req dtos.PaymentIntentRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}

	total, err := strconv.ParseFloat(req.Total, 64)
	if err != nil {
		return err
	}

	clientSecret, err := hnd.srv.CreateIntent(r.Context(), userID, total)
	if err != nil {
		return err
	}

	resp := dtos.PaymentClientSecretResponse{
		ClientSecret: string(clientSecret),
	}
	return WriteJson(w, http.StatusOK, resp)
}

func (hnd *Payment) HandleWebhookEvent(w http.ResponseWriter, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	_ = hnd.srv.ProcessWebhookEvent(
		r.Context(),
		r.Header.Get("Stripe-Signature"),
		body,
	)
	return nil
}

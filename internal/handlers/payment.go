package handlers

import (
	"io"
	"net/http"
	"strconv"

	"github.com/iypetrov/gopizza/internal/dtos"
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

func (hnd *Payment) GetPaymentMetadata(w http.ResponseWriter, r *http.Request) error {
	req, err := dtos.ParseToPaymentIntentRequest(r)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrLoadingPaymentMetadata))
		return toasts.ErrLoadingPaymentMetadata
	}

	total, err := strconv.ParseFloat(req.Total, 64)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return err
	}

	intentID, clientSecret, err := hnd.srv.GetPaymentMetadata(r.Context(), req.Email, total)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return err
	}

	resp := dtos.PaymentClientSecretResponse{
		IntentID:     intentID,
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

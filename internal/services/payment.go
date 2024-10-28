package services

import (
	"context"
	"database/sql"
	"log"

	"github.com/iypetrov/gopizza/configs"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/paymentintent"
	"github.com/stripe/stripe-go/v80/webhook"
)

type ClientSecret string

type Payment struct {
	db      *sql.DB
	queries *database.Queries
}

func NewPayment(db *sql.DB, queries *database.Queries) Payment {
	return Payment{
		db:      db,
		queries: queries,
	}
}

func (srv *Payment) GetPublishableKey() string {
	return configs.Get().Stripe.PublishableKey
}

func (srv *Payment) CreateIntent(ctx context.Context, total float64) (ClientSecret, error) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(total * 100)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		return "", err
	}

	return ClientSecret(pi.ClientSecret), nil
}

func (srv *Payment) ProcessWebhookEvent(ctx context.Context, stripeSignature string, body []byte) error {
	event, err := webhook.ConstructEvent(
		body,
		stripeSignature,
		configs.Get().Stripe.WebhookSecret,
	)
	if err != nil {
		return err
	}

	switch event.Type {
	case "payment_intent.succeeded":
		log.Println("we got the money")
	case "charge.succeeded":
		log.Println("we got the money")
	}

	return nil
}

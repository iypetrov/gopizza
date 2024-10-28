package services

import (
	"context"

	"github.com/iypetrov/gopizza/configs"
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/paymentintent"
	"github.com/stripe/stripe-go/v80/webhook"
)

type ClientSecret string

type Payment struct {
	orderSrv Order
}

func NewPayment(orderSrv Order) Payment {
	return Payment{
		orderSrv: orderSrv,
	}
}

func (srv *Payment) GetPublishableKey() string {
	return configs.Get().Stripe.PublishableKey
}

func (srv *Payment) GetPaymentMetadata(ctx context.Context, email string, total float64) (string, ClientSecret, error) {
	param := &stripe.PaymentIntentParams{
		// Customer: stripe.String(email),
		Amount:   stripe.Int64(int64(total * 100)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	intent, err := paymentintent.New(param)
	if err != nil {
		return "", "", err
	}

	return intent.ID, ClientSecret(intent.ClientSecret), nil
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
		intentID := event.Data.Object["id"].(string)
		_, err = srv.orderSrv.ChargeOrder(ctx, intentID)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

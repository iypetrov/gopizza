package services

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/configs"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/toasts"
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

func (srv *Payment) CreateIntent(ctx context.Context, userID uuid.UUID, total float64) (ClientSecret, error) {
	paramsStripe := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(total * 100)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	intent, err := paymentintent.New(paramsStripe)
	if err != nil {
		return "", err
	}

	paramsDB := database.InitOrderParams{
		ID:        uuid.New(),
		IntentID:  intent.ID,
		UserID:    uuid.NullUUID{
			UUID:  userID,
			Valid: true,
		},
		Amount:    total,
		CreatedAt: time.Now().UTC(),
	}
	_, err = srv.queries.InitOrder(ctx, paramsDB)
	if err != nil {
		return "", err
	}

	return ClientSecret(intent.ClientSecret), nil
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

		tx, err := srv.db.Begin()
		if err != nil {
			return toasts.ErrDatabaseTransactionFailed
		}
		defer tx.Rollback()

		qtx := srv.queries.WithTx(tx)
		order, err := qtx.GetOrderByIntentID(ctx, intentID)
		if err != nil {
			return err
		}

		params := database.ChargeOrderParams{
			ID: order.ID,
			UpdatedAt: sql.NullTime{
				Time:  time.Now().UTC(),
				Valid: true,
			},
		}
		_, err = qtx.ChargeOrder(ctx, params)
		if err != nil {
			return err
		}

		log.Println("order id " + order.ID.String() + " was charged")
		if order.UserID.Valid {
			log.Println("user id " + order.UserID.UUID.String() + " was charged")
			_, err = qtx.EmptyCartByUserID(ctx, uuid.NullUUID{
				UUID:  order.UserID.UUID,
				Valid: true,
			})
		}

		err = tx.Commit()
		if err != nil {
			return toasts.ErrDatabaseTransactionFailed
		}

		log.Println("order was charged")
		return nil
	}

	return nil
}

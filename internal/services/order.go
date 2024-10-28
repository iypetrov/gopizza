package services

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/toasts"
)

type Order struct {
	db      *sql.DB
	queries *database.Queries
}

func NewOrder(db *sql.DB, queries *database.Queries) Order {
	return Order{
		db:      db,
		queries: queries,
	}
}

func (srv *Order) CreateOrder(ctx context.Context, intentID string, userID uuid.UUID, total float64) (database.Order, error) {
	paramsDB := database.CreateOrderParams{
		ID:       uuid.New(),
		IntentID: intentID,
		UserID: uuid.NullUUID{
			UUID:  userID,
			Valid: true,
		},
		Amount:    total,
		CreatedAt: time.Now().UTC(),
	}
	model, err := srv.queries.CreateOrder(ctx, paramsDB)
	if err != nil {
		return database.Order{}, err
	}

	return model, nil
}

func (srv *Order) GetOrderByIntentID(ctx context.Context, intentID string) (database.GetOrderByIntentIDRow, error) {
	model, err := srv.queries.GetOrderByIntentID(ctx, intentID)
	if err != nil {
		return database.GetOrderByIntentIDRow{}, err
	}
	return model, nil
}

func (srv *Order) ChargeOrder(ctx context.Context, intentID string) (database.Order, error) {
	tx, err := srv.db.Begin()
	if err != nil {
		return database.Order{}, toasts.ErrDatabaseTransactionFailed
	}
	defer tx.Rollback()

	qtx := srv.queries.WithTx(tx)
	order, err := qtx.GetOrderByIntentID(ctx, intentID)
	if err != nil {
		return database.Order{}, err
	}

	params := database.ChargeOrderParams{
		ID: order.ID,
		UpdatedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	}
	model, err := qtx.ChargeOrder(ctx, params)
	if err != nil {
		return database.Order{}, err
	}
	log.Printf("order %s was charged\n", model.ID.String())

	if order.UserID.Valid {
		_, err = qtx.EmptyCartByUserID(ctx, uuid.NullUUID{
			UUID:  order.UserID.UUID,
			Valid: true,
		})
		log.Printf("cart of user %s was cleared\n", order.UserID.UUID.String())
	}

	err = tx.Commit()
	if err != nil {
		return database.Order{}, toasts.ErrDatabaseTransactionFailed
	}

	return model, nil
}

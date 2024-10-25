package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/toasts"
)

type Cart struct {
	db            *sql.DB
	queries       *database.Queries
}

func NewCart(db *sql.DB, queries *database.Queries) Cart {
	return Cart{
		db:            db,
		queries:       queries,
	}
}

func (srv *Cart) AddPizzaToCart(ctx context.Context, userID, pizzaID uuid.UUID) ([]database.GetCartByUserIDRow, error) {
	p := database.AddPizzaToCartParams{
		ID:        uuid.New(),
		UserID:    userID,
		PizzaID:   uuid.NullUUID{
			UUID:  pizzaID,
			Valid: true,
		},
		CreatedAt: time.Now().UTC(),
	}

	tx, err := srv.db.Begin()
	if err != nil {
		return []database.GetCartByUserIDRow{}, toasts.ErrDatabaseTransactionFailed
	}
	defer tx.Rollback()

	qtx := srv.queries.WithTx(tx)

	_, err = qtx.AddPizzaToCart(ctx, p)
	if err != nil {
		return []database.GetCartByUserIDRow{}, toasts.ErrCartItemCreation
	}

	ms, err := qtx.GetCartByUserID(ctx, p.UserID)
	if err != nil {
		return nil, toasts.ErrCartDoesNotExist
	}

	err = tx.Commit()
	if err != nil {
		return nil, toasts.ErrDatabaseTransactionFailed
	}

	return ms, nil
}

func (srv *Cart) GetCartByUserID(ctx context.Context, userID uuid.UUID) ([]database.GetCartByUserIDRow, error) {
	ms, err := srv.queries.GetCartByUserID(ctx, userID)
	if err != nil {
		return []database.GetCartByUserIDRow{}, toasts.ErrCartDoesNotExist
	}

	return ms, nil
}

func (srv *Cart) EmptyCartByUserID(ctx context.Context, userID uuid.UUID) error {
	_, err := srv.queries.EmptyCartByUserID(ctx, userID)
	if err != nil {
		return toasts.ErrCartDoesNotExist
	}

	return nil
}	

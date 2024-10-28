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
	db      *sql.DB
	queries *database.Queries
}

func NewCart(db *sql.DB, queries *database.Queries) Cart {
	return Cart{
		db:      db,
		queries: queries,
	}
}

func (srv *Cart) AddPizzaToCart(ctx context.Context, userID, pizzaID uuid.UUID) error {
	p := database.AddPizzaToCartParams{
		ID: uuid.New(),
		UserID: uuid.NullUUID{
			UUID:  userID,
			Valid: true,
		},
		PizzaID: uuid.NullUUID{
			UUID:  pizzaID,
			Valid: true,
		},
		CreatedAt: time.Now().UTC(),
	}

	_, err := srv.queries.AddPizzaToCart(ctx, p)
	if err != nil {
		return toasts.ErrCartItemCreation
	}

	return nil
}

func (srv *Cart) AddSaladToCart(ctx context.Context, userID, saladID uuid.UUID) error {
	p := database.AddSaladToCartParams{
		ID: uuid.New(),
		UserID: uuid.NullUUID{
			UUID:  userID,
			Valid: true,
		},
		SaladID: uuid.NullUUID{
			UUID:  saladID,
			Valid: true,
		},
		CreatedAt: time.Now().UTC(),
	}

	_, err := srv.queries.AddSaladToCart(ctx, p)
	if err != nil {
		return toasts.ErrCartItemCreation
	}

	return nil
}

func (srv *Cart) GetCartByUserID(ctx context.Context, userID uuid.UUID) ([]database.GetCartByUserIDRow, error) {
	ms, err := srv.queries.GetCartByUserID(ctx, uuid.NullUUID{
		UUID:  userID,
		Valid: true,
	})
	if err != nil {
		return []database.GetCartByUserIDRow{}, toasts.ErrCartDoesNotExist
	}

	return ms, nil
}

func (srv *Cart) RemoveItemFromCart(ctx context.Context, id uuid.UUID, userID uuid.UUID) ([]database.GetCartByUserIDRow, error) {
	tx, err := srv.db.Begin()
	if err != nil {
		return []database.GetCartByUserIDRow{}, toasts.ErrDatabaseTransactionFailed
	}
	defer tx.Rollback()

	qtx := srv.queries.WithTx(tx)
	err = qtx.RemoveItemFromCart(ctx, id)
	if err != nil {
		return []database.GetCartByUserIDRow{}, toasts.ErrCartDoesNotExist
	}

	ms, err := qtx.GetCartByUserID(ctx, uuid.NullUUID{
		UUID:  userID,
		Valid: true,
	})
	if err != nil {
		return []database.GetCartByUserIDRow{}, toasts.ErrCartDoesNotExist
	}

	err = tx.Commit()
	if err != nil {
		return []database.GetCartByUserIDRow{}, toasts.ErrDatabaseTransactionFailed
	}

	return ms, nil
}

func (srv *Cart) EmptyCartByUserID(ctx context.Context, userID uuid.UUID) error {
	_, err := srv.queries.EmptyCartByUserID(ctx, uuid.NullUUID{
		UUID:  userID,
		Valid: true,
	})
	if err != nil {
		return toasts.ErrCartDoesNotExist
	}

	return nil
}

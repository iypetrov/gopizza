package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/toasts"
	"github.com/lib/pq"
)

type Pizza struct {
	db      *sql.DB
	queries *database.Queries
}

func NewPizza(db *sql.DB, queries *database.Queries) Pizza {
	return Pizza{
		db:      db,
		queries: queries,
	}
}

func (srv *Pizza) CreatePizza(ctx context.Context, p database.CreatePizzaParams) (database.Pizza, error) {
	p.ID = uuid.New()
	p.UpdatedAt = time.Now()

	m, err := srv.queries.CreatePizza(ctx, p)
	if err != nil {
		var pgErr *pq.Error

		ok := errors.As(err, &pgErr)
		if ok {
			if pgErr.Code == "23505" {
				return database.Pizza{}, toasts.ErrPizzasAlreadyExists
			}
		}

		return database.Pizza{}, toasts.ErrPizzaCreation
	}

	return m, nil
}

func (srv *Pizza) GetPizzaByID(ctx context.Context, id uuid.UUID) (database.Pizza, error) {
	m, err := srv.queries.GetPizzaByID(ctx, id)
	if err != nil {
		return database.Pizza{}, toasts.ErrPizzaNotFound
	}

	return m, nil
}

func (srv *Pizza) GetAllPizzas(ctx context.Context) ([]database.Pizza, error) {
	ms, err := srv.queries.GetAllPizzas(ctx)
	if err != nil {
		return nil, toasts.ErrPizzaFailedToLoad
	}

	return ms, nil
}

func (srv *Pizza) UpdateModel(ctx context.Context, id uuid.UUID, p database.UpdatePizzaParams) (database.Pizza, error) {
	p.ID = id

	m, err := srv.queries.UpdatePizza(ctx, p)
	if err != nil {
		var pgErr *pq.Error

		ok := errors.As(err, &pgErr)
		if ok {
			if pgErr.Code == "23505" {
				return database.Pizza{}, toasts.ErrPizzasAlreadyExists
			}
		}

		return database.Pizza{}, toasts.ErrPizzaUpdating
	}

	return m, nil
}

func (srv *Pizza) DeletePizzaByID(ctx context.Context, id uuid.UUID) ([]database.Pizza, error) {
	tx, err := srv.db.Begin()
	if err != nil {
		return []database.Pizza{}, toasts.ErrDatabaseTransactionFailed
	}
	defer tx.Rollback()

	qtx := srv.queries.WithTx(tx)

	_, err = qtx.DeletePizzaByID(ctx, id)
	if err != nil {
		var pgErr *pq.Error

		ok := errors.As(err, &pgErr)
		if ok {
			if pgErr.Code == "23503" {
				return []database.Pizza{}, toasts.ErrPizzaNotFound
			}
		}

		return []database.Pizza{}, toasts.ErrPizzaDeletion
	}

	ms, err := qtx.GetAllPizzas(ctx)
	if err != nil {
		return nil, toasts.ErrPizzaFailedToLoad
	}

	err = tx.Commit()
	if err != nil {
		return []database.Pizza{}, toasts.ErrDatabaseTransactionFailed
	}

	return ms, nil
}

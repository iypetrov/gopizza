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

type Salad struct {
	db      *sql.DB
	queries *database.Queries
}

func NewSalad(db *sql.DB, queries *database.Queries) Salad {
	return Salad{
		db:      db,
		queries: queries,
	}
}

func (srv *Salad) CreateSalad(ctx context.Context, p database.CreateSaladParams) ([]database.Salad, error) {
	p.ID = uuid.New()
	p.UpdatedAt = time.Now().UTC()
	p.CreatedAt = time.Now().UTC()

	tx, err := srv.db.Begin()
	if err != nil {
		return []database.Salad{}, toasts.ErrDatabaseTransactionFailed
	}
	defer tx.Rollback()

	qtx := srv.queries.WithTx(tx)

	_, err = qtx.CreateSalad(ctx, p)
	if err != nil {
		var pgErr *pq.Error

		ok := errors.As(err, &pgErr)
		if ok {
			if pgErr.Code == "23505" {
				return []database.Salad{}, toasts.ErrSaladsAlreadyExists
			}
		}

		return nil, toasts.ErrSaladCreation
	}

	ms, err := qtx.GetAllSalads(ctx)
	if err != nil {
		return []database.Salad{}, toasts.ErrSaladFailedToLoad
	}

	err = tx.Commit()
	if err != nil {
		return []database.Salad{}, toasts.ErrDatabaseTransactionFailed
	}

	return ms, nil
}

func (srv *Salad) GetSaladByID(ctx context.Context, id uuid.UUID) (database.Salad, error) {
	m, err := srv.queries.GetSaladByID(ctx, id)
	if err != nil {
		return database.Salad{}, toasts.ErrSaladNotFound
	}

	return m, nil
}

func (srv *Salad) GetAllSalads(ctx context.Context) ([]database.Salad, error) {
	ms, err := srv.queries.GetAllSalads(ctx)
	if err != nil {
		return []database.Salad{}, toasts.ErrSaladFailedToLoad
	}

	return ms, nil
}

func (srv *Salad) UpdateModel(ctx context.Context, id uuid.UUID, p database.UpdateSaladParams) (database.Salad, error) {
	p.ID = id

	m, err := srv.queries.UpdateSalad(ctx, p)
	if err != nil {
		var pgErr *pq.Error

		ok := errors.As(err, &pgErr)
		if ok {
			if pgErr.Code == "23505" {
				return database.Salad{}, toasts.ErrSaladsAlreadyExists
			}
		}

		return database.Salad{}, toasts.ErrSaladUpdating
	}

	return m, nil
}

func (srv *Salad) DeleteSaladByID(ctx context.Context, id uuid.UUID) ([]database.Salad, error) {
	tx, err := srv.db.Begin()
	if err != nil {
		return []database.Salad{}, toasts.ErrDatabaseTransactionFailed
	}
	defer tx.Rollback()

	qtx := srv.queries.WithTx(tx)

	_, err = qtx.DeleteSaladByID(ctx, id)
	if err != nil {
		var pgErr *pq.Error

		ok := errors.As(err, &pgErr)
		if ok {
			if pgErr.Code == "23503" {
				return []database.Salad{}, toasts.ErrSaladNotFound
			}
		}

		return []database.Salad{}, toasts.ErrSaladDeletion
	}

	ms, err := qtx.GetAllSalads(ctx)
	if err != nil {
		return []database.Salad{}, toasts.ErrSaladFailedToLoad
	}

	err = tx.Commit()
	if err != nil {
		return []database.Salad{}, toasts.ErrDatabaseTransactionFailed
	}

	return ms, nil
}

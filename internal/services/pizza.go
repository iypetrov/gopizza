package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/toasts"
	"github.com/lib/pq"
	"time"
)

type Pizza interface {
	CreatePizza(ctx context.Context, p database.CreatePizzaParams) (database.Pizza, error)
	GetPizzaByID(ctx context.Context, id uuid.UUID) (database.Pizza, error)
	GetAllPizzas(ctx context.Context, p database.GetAllPizzasParams) ([]database.Pizza, error)
	UpdateModel(ctx context.Context, id uuid.UUID, p database.UpdatePizzaParams) (database.Pizza, error)
	DeletePizzaByID(ctx context.Context, id uuid.UUID) (database.Pizza, error)
}

type PizzaImpl struct {
	ctx context.Context
	db  *database.Queries
}

func NewPizza(ctx context.Context, db *database.Queries) *PizzaImpl {
	return &PizzaImpl{
		ctx: ctx,
		db:  db,
	}
}

func (srv *PizzaImpl) CreatePizza(ctx context.Context, p database.CreatePizzaParams) (database.Pizza, error) {
	p.ID = uuid.New()
	p.UpdatedAt = time.Now()

	m, err := srv.db.CreatePizza(ctx, p)
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

func (srv *PizzaImpl) GetPizzaByID(ctx context.Context, id uuid.UUID) (database.Pizza, error) {
	m, err := srv.db.GetPizzaByID(ctx, id)
	if err != nil {
		return database.Pizza{}, toasts.ErrPizzaNotFound
	}

	return m, nil
}

func (srv *PizzaImpl) GetAllPizzas(ctx context.Context, p database.GetAllPizzasParams) ([]database.Pizza, error) {
	ms, err := srv.db.GetAllPizzas(ctx, p)
	if err != nil {
		return nil, toasts.ErrPizzaFailedToLoad
	}

	return ms, nil
}

func (srv *PizzaImpl) UpdateModel(ctx context.Context, id uuid.UUID, p database.UpdatePizzaParams) (database.Pizza, error) {
	p.ID = id

	m, err := srv.db.UpdatePizza(ctx, p)
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

func (srv *PizzaImpl) DeletePizzaByID(ctx context.Context, id uuid.UUID) (database.Pizza, error) {
	m, err := srv.db.DeletePizzaByID(ctx, id)
	if err != nil {
		var pgErr *pq.Error

		ok := errors.As(err, &pgErr)
		if ok {
			if pgErr.Code == "23503" {
				return database.Pizza{}, toasts.ErrPizzaNotFound
			}
		}

		return database.Pizza{}, toasts.ErrPizzaDeletion
	}

	return m, nil
}
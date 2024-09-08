package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/mapper"
	"github.com/iypetrov/gopizza/internal/model"
	"github.com/iypetrov/gopizza/internal/repository"
	"github.com/iypetrov/gopizza/internal/toast"
	"github.com/lib/pq"
)

type Pizza interface {
	CreatePizza(ctx context.Context, m model.Pizza) (model.Pizza, error)
	GetPizzaByID(ctx context.Context, id uuid.UUID) (model.Pizza, error)
	GetAllPizzas(ctx context.Context, lastID uuid.UUID, lastPrice float64, pageSize int32) ([]model.Pizza, error)
	UpdateModel(ctx context.Context, id uuid.UUID, m model.Pizza) (model.Pizza, error)
	DeletePizzaByID(ctx context.Context, id uuid.UUID) (model.Pizza, error)
}

type PizzaImpl struct {
	ctx context.Context
	rep repository.Pizza
}

func NewPizza(ctx context.Context, repository repository.Pizza) *PizzaImpl {
	return &PizzaImpl{
		ctx: ctx,
		rep: repository,
	}
}

func (srv *PizzaImpl) CreatePizza(ctx context.Context, m model.Pizza) (model.Pizza, error) {
	//err := m.Validate()
	//if err != nil {
	//	return model.Pizza{}, err
	//}

	e, err := srv.rep.CreatePizza(ctx, m)
	if err != nil {
		var pgErr *pq.Error

		ok := errors.As(err, &pgErr)
		if ok {
			if pgErr.Code == "23505" {
				return model.Pizza{}, toast.ErrPizzasAlreadyExists
			}
		}

		return model.Pizza{}, toast.ErrPizzaCreation
	}

	return mapper.PizzaEntityToModel(e), nil
}

func (srv *PizzaImpl) GetPizzaByID(ctx context.Context, id uuid.UUID) (model.Pizza, error) {
	e, err := srv.rep.GetPizzaByID(ctx, id)
	if err != nil {
		return model.Pizza{}, toast.ErrPizzaNotFound
	}

	return mapper.PizzaEntityToModel(e), nil
}

func (srv *PizzaImpl) GetAllPizzas(ctx context.Context, lastID uuid.UUID, lastPrice float64, pageSize int32) ([]model.Pizza, error) {
	es, err := srv.rep.GetAllPizzas(ctx, lastID, lastPrice, pageSize)
	if err != nil {
		return nil, toast.ErrPizzaFailedToLoad
	}

	var ms []model.Pizza
	for _, e := range es {
		ms = append(ms, mapper.PizzaEntityToModel(e))
	}

	return ms, nil
}

func (srv *PizzaImpl) UpdateModel(ctx context.Context, id uuid.UUID, m model.Pizza) (model.Pizza, error) {
	m.ID = id

	//err := m.Validate()
	//if err != nil {
	//	return model.Pizza{}, err
	//}

	e, err := srv.rep.UpdatePizza(ctx, m)
	if err != nil {
		var pgErr *pq.Error

		ok := errors.As(err, &pgErr)
		if ok {
			if pgErr.Code == "23505" {
				return model.Pizza{}, toast.ErrPizzasAlreadyExists
			}
		}

		return model.Pizza{}, toast.ErrPizzaUpdating
	}

	return mapper.PizzaEntityToModel(e), nil
}

func (srv *PizzaImpl) DeletePizzaByID(ctx context.Context, id uuid.UUID) (model.Pizza, error) {
	e, err := srv.rep.DeletePizzaByID(ctx, id)
	if err != nil {
		var pgErr *pq.Error

		ok := errors.As(err, &pgErr)
		if ok {
			if pgErr.Code == "23503" {
				return model.Pizza{}, toast.ErrPizzaNotFound
			}
		}

		return model.Pizza{}, toast.ErrPizzaDeletion
	}

	return mapper.PizzaEntityToModel(e), nil
}

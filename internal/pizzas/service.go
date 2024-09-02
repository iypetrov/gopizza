package pizzas

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/utils"
	"github.com/lib/pq"
)

type PizzaService interface {
	CreatePizzaModel(ctx context.Context, model PizzaModel) (PizzaModel, error)
	GetPizzaModelByID(ctx context.Context, id uuid.UUID) (PizzaModel, error)
	UpdatePizzaModel(ctx context.Context, id uuid.UUID, model PizzaModel) (PizzaModel, error)
	DeletePizzaModelByID(ctx context.Context, id uuid.UUID) (PizzaModel, error)
}

type Service struct {
	ctx        context.Context
	repository PizzaRepository
}

func NewService(ctx context.Context, repository PizzaRepository) *Service {
	return &Service{
		ctx:        ctx,
		repository: repository,
	}
}

func (srv *Service) CreatePizzaModel(ctx context.Context, model PizzaModel) (PizzaModel, error) {
	err := model.Validate()
	if err != nil {
		return PizzaModel{}, err
	}

	entity, err := srv.repository.CreatePizzaEntity(ctx, model)
	if err != nil {
		var pgErr *pq.Error
		ok := errors.As(err, &pgErr)
		if ok {
			if pgErr.Code == "23505" {
				return PizzaModel{}, utils.BadRequest(ErrPizzasAlreadyExists)
			}
		}
		return PizzaModel{}, utils.InternalServerError(ErrCreatingPizza)
	}
	return entity.ToModel(), nil
}

func (srv *Service) GetPizzaModelByID(ctx context.Context, id uuid.UUID) (PizzaModel, error) {
	entity, err := srv.repository.GetPizzaEntityByID(ctx, id)
	if err != nil {
		return PizzaModel{}, utils.NotFound(ErrPizzaNotFound)
	}
	return entity.ToModel(), nil
}

func (srv *Service) UpdatePizzaModel(ctx context.Context, id uuid.UUID, model PizzaModel) (PizzaModel, error) {
	model.ID = id

	err := model.Validate()
	if err != nil {
		return PizzaModel{}, err
	}

	entity, err := srv.repository.UpdatePizzaEntity(ctx, model)
	if err != nil {
		var pgErr *pq.Error
		ok := errors.As(err, &pgErr)
		if ok {
			if pgErr.Code == "23505" {
				return PizzaModel{}, utils.BadRequest(ErrPizzasAlreadyExists)
			}
		}
		return PizzaModel{}, utils.InternalServerError(ErrUpdatingPizza)
	}
	return entity.ToModel(), nil
}

func (srv *Service) DeletePizzaModelByID(ctx context.Context, id uuid.UUID) (PizzaModel, error) {
	entity, err := srv.repository.DeletePizzaEntityByID(ctx, id)
	if err != nil {
		var pgErr *pq.Error
		ok := errors.As(err, &pgErr)
		fmt.Println(pgErr)
		if ok {
			if pgErr.Code == "23503" {
				return PizzaModel{}, utils.BadRequest(ErrPizzaNotFound)
			}
		}
		return PizzaModel{}, utils.InternalServerError(ErrDeletingPizza)
	}

	return entity.ToModel(), nil
}

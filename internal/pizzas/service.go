package pizzas

import (
	"context"
	"errors"
	"github.com/iypetrov/gopizza/internal/config"
	"github.com/iypetrov/gopizza/internal/utils"
	"github.com/lib/pq"
)

type PizzaService interface {
	CreatePizzaModel(ctx context.Context, model PizzaModel) (PizzaModel, error)
}

type Service struct {
	ctx        context.Context
	log        *config.Logger
	repository PizzaRepository
}

func NewService(ctx context.Context, log *config.Logger, repository PizzaRepository) *Service {
	return &Service{
		ctx:        ctx,
		log:        log,
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

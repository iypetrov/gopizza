package pizzas

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/config"
	"github.com/iypetrov/gopizza/internal/utils"
	"github.com/lib/pq"
)

type PizzaService interface {
	CreatePizzaModel(ctx context.Context, model PizzaModel) (PizzaModel, error)
	GetPizzaModelByID(ctx context.Context, id string) (PizzaModel, error)
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

func (srv *Service) GetPizzaModelByID(ctx context.Context, id string) (PizzaModel, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return PizzaModel{}, utils.InvalidUUID()
	}

	entity, err := srv.repository.GetPizzaEntityByID(ctx, uuidID)
	if err != nil {
		return PizzaModel{}, utils.NotFound(ErrPizzaNotFound)
	}
	return entity.ToModel(), nil
}

package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/entity"
	"github.com/iypetrov/gopizza/internal/mapper"
	"github.com/iypetrov/gopizza/internal/model"
)

type Pizza interface {
	CreatePizza(ctx context.Context, m model.Pizza) (entity.Pizza, error)
	GetPizzaByID(ctx context.Context, id uuid.UUID) (entity.Pizza, error)
	GetAllPizzas(ctx context.Context, lastID uuid.UUID, lastPrice float64, pageSize int32) ([]entity.Pizza, error)
	UpdatePizza(ctx context.Context, m model.Pizza) (entity.Pizza, error)
	DeletePizzaByID(ctx context.Context, id uuid.UUID) (entity.Pizza, error)
}

type PizzaImpl struct {
	db *database.Queries
}

func NewPizza(db *database.Queries) *PizzaImpl {
	return &PizzaImpl{db: db}
}

func (rep *PizzaImpl) CreatePizza(ctx context.Context, m model.Pizza) (entity.Pizza, error) {
	p, err := rep.db.CreatePizza(ctx, database.CreatePizzaParams{
		ID:         m.ID,
		Name:       m.Name,
		Tomatoes:   m.Tomatoes,
		Garlic:     m.Garlic,
		Onion:      m.Onion,
		Parmesan:   m.Parmesan,
		Cheddar:    m.Cheddar,
		Pepperoni:  m.Pepperoni,
		Sausage:    m.Sausage,
		Ham:        m.Ham,
		Bacon:      m.Bacon,
		Chicken:    m.Chicken,
		Salami:     m.Salami,
		GroundBeef: m.GroundBeef,
		Mushrooms:  m.Mushrooms,
		Olives:     m.Olives,
		Spinach:    m.Spinach,
		Pineapple:  m.Pineapple,
		Arugula:    m.Arugula,
		Anchovies:  m.Anchovies,
		Capers:     m.Capers,
		ImageUrl:   m.ImageUrl,
		Price:      m.Price,
		UpdatedAt:  m.UpdatedAt,
	})
	return mapper.PizzaEntityFromSQLC(p), err
}

func (rep *PizzaImpl) GetPizzaByID(ctx context.Context, id uuid.UUID) (entity.Pizza, error) {
	pizza, err := rep.db.GetPizzaByID(ctx, id)
	return mapper.PizzaEntityFromSQLC(pizza), err
}

func (rep *PizzaImpl) GetAllPizzas(ctx context.Context, lastID uuid.UUID, lastPrice float64, pageSize int32) ([]entity.Pizza, error) {
	var es []entity.Pizza
	ps, err := rep.db.GetAllPizzas(ctx, database.GetAllPizzasParams{
		ID:       lastID,
		Price:    lastPrice,
		PageSize: pageSize,
	})
	for _, p := range ps {
		es = append(es, mapper.PizzaEntityFromSQLC(p))
	}
	return es, err
}

func (rep *PizzaImpl) UpdatePizza(ctx context.Context, m model.Pizza) (entity.Pizza, error) {
	p, err := rep.db.UpdatePizza(ctx, database.UpdatePizzaParams{
		ID:         m.ID,
		Name:       m.Name,
		Tomatoes:   m.Tomatoes,
		Garlic:     m.Garlic,
		Onion:      m.Onion,
		Parmesan:   m.Parmesan,
		Cheddar:    m.Cheddar,
		Pepperoni:  m.Pepperoni,
		Sausage:    m.Sausage,
		Ham:        m.Ham,
		Bacon:      m.Bacon,
		Chicken:    m.Chicken,
		Salami:     m.Salami,
		GroundBeef: m.GroundBeef,
		Mushrooms:  m.Mushrooms,
		Olives:     m.Olives,
		Spinach:    m.Spinach,
		Pineapple:  m.Pineapple,
		Arugula:    m.Arugula,
		Anchovies:  m.Anchovies,
		Capers:     m.Capers,
		ImageUrl:   m.ImageUrl,
		Price:      m.Price,
		UpdatedAt:  m.UpdatedAt,
	})
	return mapper.PizzaEntityFromSQLC(p), err
}

func (rep *PizzaImpl) DeletePizzaByID(ctx context.Context, id uuid.UUID) (entity.Pizza, error) {
	p, err := rep.db.DeletePizzaByID(ctx, id)
	return mapper.PizzaEntityFromSQLC(p), err
}

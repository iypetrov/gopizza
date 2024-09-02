package pizzas

import (
	"context"
	"github.com/iypetrov/gopizza/internal/database"
)

type PizzaRepository interface {
	CreatePizzaEntity(ctx context.Context, model PizzaModel) (PizzaEntity, error)
}

type Repository struct {
	db *database.Queries
}

func NewRepository(db *database.Queries) *Repository {
	return &Repository{db: db}
}

func (rep *Repository) CreatePizzaEntity(ctx context.Context, model PizzaModel) (PizzaEntity, error) {
	var entity PizzaEntity
	pizza, err := rep.db.CreatePizza(ctx, database.CreatePizzaParams{
		ID:         model.ID,
		Name:       model.Name,
		Tomatoes:   model.Tomatoes,
		Garlic:     model.Garlic,
		Onion:      model.Onion,
		Parmesan:   model.Parmesan,
		Cheddar:    model.Cheddar,
		Pepperoni:  model.Pepperoni,
		Sausage:    model.Sausage,
		Ham:        model.Ham,
		Bacon:      model.Bacon,
		Chicken:    model.Chicken,
		Salami:     model.Salami,
		GroundBeef: model.GroundBeef,
		Mushrooms:  model.Mushrooms,
		Olives:     model.Olives,
		Spinach:    model.Spinach,
		Pineapple:  model.Pineapple,
		Arugula:    model.Arugula,
		Anchovies:  model.Anchovies,
		Capers:     model.Capers,
		ImageUrl:   model.ImageUrl,
		Price:      model.Price,
		UpdatedAt:  model.UpdatedAt,
	})
	return entity.FromSqlC(pizza), err
}

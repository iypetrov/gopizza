package pizzas

import (
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/database"
	"time"
)

type PizzaEntity struct {
	ID         uuid.UUID
	Name       string
	Tomatoes   bool
	Garlic     bool
	Onion      bool
	Parmesan   bool
	Cheddar    bool
	Pepperoni  bool
	Sausage    bool
	Ham        bool
	Bacon      bool
	Chicken    bool
	Salami     bool
	GroundBeef bool
	Mushrooms  bool
	Olives     bool
	Spinach    bool
	Pineapple  bool
	Arugula    bool
	Anchovies  bool
	Capers     bool
	ImageUrl   string
	Price      float64
	UpdatedAt  time.Time
}

func (e *PizzaEntity) FromSqlC(pizza database.Pizza) PizzaEntity {
	e.ID = pizza.ID
	e.Name = pizza.Name
	e.Tomatoes = pizza.Tomatoes
	e.Garlic = pizza.Garlic
	e.Onion = pizza.Onion
	e.Parmesan = pizza.Parmesan
	e.Cheddar = pizza.Cheddar
	e.Pepperoni = pizza.Pepperoni
	e.Sausage = pizza.Sausage
	e.Ham = pizza.Ham
	e.Bacon = pizza.Bacon
	e.Chicken = pizza.Chicken
	e.Salami = pizza.Salami
	e.GroundBeef = pizza.GroundBeef
	e.Mushrooms = pizza.Mushrooms
	e.Olives = pizza.Olives
	e.Spinach = pizza.Spinach
	e.Pineapple = pizza.Pineapple
	e.Arugula = pizza.Arugula
	e.Anchovies = pizza.Anchovies
	e.Capers = pizza.Capers
	e.ImageUrl = pizza.ImageUrl
	e.Price = pizza.Price
	e.UpdatedAt = pizza.UpdatedAt

	return *e
}

func (e *PizzaEntity) ToModel() PizzaModel {
	return PizzaModel{
		ID:         e.ID,
		Name:       e.Name,
		Tomatoes:   e.Tomatoes,
		Garlic:     e.Garlic,
		Onion:      e.Onion,
		Parmesan:   e.Parmesan,
		Cheddar:    e.Cheddar,
		Pepperoni:  e.Pepperoni,
		Sausage:    e.Sausage,
		Ham:        e.Ham,
		Bacon:      e.Bacon,
		Chicken:    e.Chicken,
		Salami:     e.Salami,
		GroundBeef: e.GroundBeef,
		Mushrooms:  e.Mushrooms,
		Olives:     e.Olives,
		Spinach:    e.Spinach,
		Pineapple:  e.Pineapple,
		Arugula:    e.Arugula,
		Anchovies:  e.Anchovies,
		Capers:     e.Capers,
		ImageUrl:   e.ImageUrl,
		Price:      e.Price,
		UpdatedAt:  e.UpdatedAt,
	}
}

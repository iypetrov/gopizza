package pizzas

import (
	"github.com/google/uuid"
	"time"
)

type CreatePizzaRequestDto struct {
	Name       string  `json:"name"`
	Tomatoes   bool    `json:"tomatoes"`
	Garlic     bool    `json:"garlic"`
	Onion      bool    `json:"onion"`
	Parmesan   bool    `json:"parmesan"`
	Cheddar    bool    `json:"cheddar"`
	Pepperoni  bool    `json:"pepperoni"`
	Sausage    bool    `json:"sausage"`
	Ham        bool    `json:"ham"`
	Bacon      bool    `json:"bacon"`
	Chicken    bool    `json:"chicken"`
	Salami     bool    `json:"salami"`
	GroundBeef bool    `json:"groundBeef"`
	Mushrooms  bool    `json:"mushrooms"`
	Olives     bool    `json:"olives"`
	Spinach    bool    `json:"spinach"`
	Pineapple  bool    `json:"pineapple"`
	Arugula    bool    `json:"arugula"`
	Anchovies  bool    `json:"anchovies"`
	Capers     bool    `json:"capers"`
	ImageUrl   string  `json:"imageUrl"`
	Price      float64 `json:"price"`
}

func (d *CreatePizzaRequestDto) ToModel() PizzaModel {
	return PizzaModel{
		ID:         uuid.New(),
		Name:       d.Name,
		Tomatoes:   d.Tomatoes,
		Garlic:     d.Garlic,
		Onion:      d.Onion,
		Parmesan:   d.Parmesan,
		Cheddar:    d.Cheddar,
		Pepperoni:  d.Pepperoni,
		Sausage:    d.Sausage,
		Ham:        d.Ham,
		Bacon:      d.Bacon,
		Chicken:    d.Chicken,
		Salami:     d.Salami,
		GroundBeef: d.GroundBeef,
		Mushrooms:  d.Mushrooms,
		Olives:     d.Olives,
		Spinach:    d.Spinach,
		Pineapple:  d.Pineapple,
		Arugula:    d.Arugula,
		Anchovies:  d.Anchovies,
		Capers:     d.Capers,
		ImageUrl:   d.ImageUrl,
		Price:      d.Price,
		UpdatedAt:  time.Now(),
	}
}

type PizzaResponseDto struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Tomatoes   bool      `json:"tomatoes"`
	Garlic     bool      `json:"garlic"`
	Onion      bool      `json:"onion"`
	Parmesan   bool      `json:"parmesan"`
	Cheddar    bool      `json:"cheddar"`
	Pepperoni  bool      `json:"pepperoni"`
	Sausage    bool      `json:"sausage"`
	Ham        bool      `json:"ham"`
	Bacon      bool      `json:"bacon"`
	Chicken    bool      `json:"chicken"`
	Salami     bool      `json:"salami"`
	GroundBeef bool      `json:"groundBeef"`
	Mushrooms  bool      `json:"mushrooms"`
	Olives     bool      `json:"olives"`
	Spinach    bool      `json:"spinach"`
	Pineapple  bool      `json:"pineapple"`
	Arugula    bool      `json:"arugula"`
	Anchovies  bool      `json:"anchovies"`
	Capers     bool      `json:"capers"`
	ImageUrl   string    `json:"imageUrl"`
	Price      float64   `json:"price"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

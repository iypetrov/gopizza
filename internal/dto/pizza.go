package dto

import (
	"github.com/google/uuid"
	"time"
)

type PizzaRequest struct {
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

type PizzaResponse struct {
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

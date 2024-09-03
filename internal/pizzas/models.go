package pizzas

import (
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/utils"
	"time"
)

type PizzaModel struct {
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

func (m *PizzaModel) Validate() error {
	var errors []error
	if len(m.Name) == 0 {
		errors = append(errors, ErrNameIsRequired)
	}
	if len(m.ImageUrl) == 0 {
		errors = append(errors, ErrImageUrlIsRequired)
	}
	if m.Price <= 0 {
		errors = append(errors, ErrPriceShouldBePositiveNumber)
	}

	if len(errors) > 0 {
		return utils.InvalidRequestData(errors)
	}

	return nil
}

func (m *PizzaModel) ToEntity() PizzaEntity {
	return PizzaEntity{
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
	}
}

func (m *PizzaModel) ToDto() PizzaResponseDto {
	return PizzaResponseDto{
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
	}
}

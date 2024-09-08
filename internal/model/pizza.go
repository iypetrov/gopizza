package model

import (
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/toast"
	"time"
)

type Pizza struct {
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

func (m *Pizza) Validate() []error {
	var errs []error
	if len(m.Name) == 0 {
		errs = append(errs, toast.ErrNameIsRequired)
	}
	if len(m.ImageUrl) == 0 {
		errs = append(errs, toast.ErrImageUrlIsRequired)
	}
	if m.Price <= 0 {
		errs = append(errs, toast.ErrPriceShouldBePositiveNumber)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

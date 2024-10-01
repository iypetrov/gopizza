package dtos

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/toasts"
)

type PizzaRequest struct {
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
}

func (req *PizzaRequest) Validate() map[string]string {
	errs := make(map[string]string)

	if len(req.Name) == 0 {
		errs["name"] = toasts.ErrNameIsRequired.Error()
	}

	if len(req.ImageUrl) == 0 {
		errs["imageUrl"] = toasts.ErrImageUrlIsRequired.Error()
	}

	if req.Price == 0 {
		errs["price"] = toasts.ErrPriceShouldBePositiveNumber.Error()
	}

	return errs
}

func ParseToPizzaRequest(r *http.Request) (PizzaRequest, error) {
	err := r.ParseForm()
	if err != nil {
		return PizzaRequest{}, err
	}

	var req PizzaRequest
	req.Name = parseString(r, "name")
	req.Tomatoes = parseBool(r, "tomatoes")
	req.Garlic = parseBool(r, "garlic")
	req.Onion = parseBool(r, "onion")
	req.Parmesan = parseBool(r, "parmesan")
	req.Cheddar = parseBool(r, "cheddar")
	req.Pepperoni = parseBool(r, "pepperoni")
	req.Sausage = parseBool(r, "sausage")
	req.Ham = parseBool(r, "ham")
	req.Bacon = parseBool(r, "bacon")
	req.Chicken = parseBool(r, "chicken")
	req.Salami = parseBool(r, "salami")
	req.GroundBeef = parseBool(r, "groundBeef")
	req.Mushrooms = parseBool(r, "mushrooms")
	req.Olives = parseBool(r, "olives")
	req.Spinach = parseBool(r, "spinach")
	req.Pineapple = parseBool(r, "pineapple")
	req.Arugula = parseBool(r, "arugula")
	req.Anchovies = parseBool(r, "anchovies")
	req.Capers = parseBool(r, "capers")
	req.ImageUrl = parseString(r, "imageUrl")
	req.Price = parseFloat(r, "price")

	return req, nil
}

type PizzaResponse struct {
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

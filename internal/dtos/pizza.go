package dtos

import (
	"io"
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
	Image      io.Reader
	Price      float64
}

func (req *PizzaRequest) Validate() map[string]string {
	errs := make(map[string]string)

	if len(req.Name) == 0 {
		errs["name"] = toasts.ErrNameIsRequired.Error()
	}

	if req.Image == nil {
		errs["image"] = toasts.ErrImageIsRequired.Error()
	} else {
		buf := make([]byte, 1)
		if _, err := req.Image.Read(buf); err == io.EOF {
			errs["image"] = toasts.ErrImageIsRequired.Error()
		} else {
			if seeker, ok := req.Image.(io.Seeker); ok {
				seeker.Seek(0, io.SeekStart)
			}
		}
	}

	if req.Price == 0 {
		errs["price"] = toasts.ErrPriceShouldBePositiveNumber.Error()
	}

	return errs
}

func ParseToPizzaRequest(r *http.Request) (PizzaRequest, error) {
	// 10 MB limit
	err := r.ParseMultipartForm(10 << 20)
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
	req.Price = parseFloat(r, "price")
	file, _, err := r.FormFile("image")
	if err != nil {
		return PizzaRequest{}, err
	}
	defer file.Close()
	req.Image = file

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

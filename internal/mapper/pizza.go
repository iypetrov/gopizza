package mapper

import (
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/dto"
	"github.com/iypetrov/gopizza/internal/entity"
	"github.com/iypetrov/gopizza/internal/model"
	"time"
)

func PizzaRequestToModel(req dto.PizzaRequest, id uuid.UUID, updateAt time.Time) model.Pizza {
	return model.Pizza{
		ID:         id,
		Name:       req.Name,
		Tomatoes:   req.Tomatoes,
		Garlic:     req.Garlic,
		Onion:      req.Onion,
		Parmesan:   req.Parmesan,
		Cheddar:    req.Cheddar,
		Pepperoni:  req.Pepperoni,
		Sausage:    req.Sausage,
		Ham:        req.Ham,
		Bacon:      req.Bacon,
		Chicken:    req.Chicken,
		Salami:     req.Salami,
		GroundBeef: req.GroundBeef,
		Mushrooms:  req.Mushrooms,
		Olives:     req.Olives,
		Spinach:    req.Spinach,
		Pineapple:  req.Pineapple,
		Arugula:    req.Arugula,
		Anchovies:  req.Anchovies,
		Capers:     req.Capers,
		ImageUrl:   req.ImageUrl,
		Price:      req.Price,
		UpdatedAt:  updateAt,
	}
}

func PizzaModelToResponse(m model.Pizza) dto.PizzaResponse {
	return dto.PizzaResponse{
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

func PizzaModelToEntity(m model.Pizza) entity.Pizza {
	return entity.Pizza{
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

func PizzaEntityToModel(e entity.Pizza) model.Pizza {
	return model.Pizza{
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

func PizzaEntityFromSQLC(p database.Pizza) entity.Pizza {
	return entity.Pizza{
		ID:         p.ID,
		Name:       p.Name,
		Tomatoes:   p.Tomatoes,
		Garlic:     p.Garlic,
		Onion:      p.Onion,
		Parmesan:   p.Parmesan,
		Cheddar:    p.Cheddar,
		Pepperoni:  p.Pepperoni,
		Sausage:    p.Sausage,
		Ham:        p.Ham,
		Bacon:      p.Bacon,
		Chicken:    p.Chicken,
		Salami:     p.Salami,
		GroundBeef: p.GroundBeef,
		Mushrooms:  p.Mushrooms,
		Olives:     p.Olives,
		Spinach:    p.Spinach,
		Pineapple:  p.Pineapple,
		Arugula:    p.Arugula,
		Anchovies:  p.Anchovies,
		Capers:     p.Capers,
		ImageUrl:   p.ImageUrl,
		Price:      p.Price,
		UpdatedAt:  p.UpdatedAt,
	}
}

package pizzas

import (
	"fmt"
)

// TODO: Add localizations for the error messages
var (
	ErrNameIsRequired              = fmt.Errorf("name is required")
	ErrImageUrlIsRequired          = fmt.Errorf("image url is required")
	ErrPriceShouldBePositiveNumber = fmt.Errorf("price should be positive number")
	ErrPizzasAlreadyExists         = fmt.Errorf("pizza with this name already exists")
	ErrCreatingPizza               = fmt.Errorf("internal server error, failed to create a pizza")
	ErrPizzaNotFound               = fmt.Errorf("pizza not found")
	ErrUpdatingPizza               = fmt.Errorf("internal server error, failed to update a pizza")
	ErrDeletingPizza               = fmt.Errorf("internal server error, failed to delete a pizza")
)

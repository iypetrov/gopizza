package pizzas

import (
	"fmt"
)

var (
	ErrNameIsRequired              = fmt.Errorf("name is required")
	ErrImageUrlIsRequired          = fmt.Errorf("image url is required")
	ErrPriceShouldBePositiveNumber = fmt.Errorf("price should be positive number")
	ErrNotValidQueryParams         = fmt.Errorf("not valid query params")
	ErrPizzasAlreadyExists         = fmt.Errorf("pizza with this name already exists")
	ErrPizzaCreation               = fmt.Errorf("internal server error, failed to create a pizza")
	ErrPizzaNotFound               = fmt.Errorf("pizza not found")
	ErrPizzaFailedToLoad           = fmt.Errorf("pizza failed to load")
	ErrPizzaUpdating               = fmt.Errorf("internal server error, failed to update a pizza")
	ErrPizzaDeletion               = fmt.Errorf("internal server error, failed to delete a pizza")
)

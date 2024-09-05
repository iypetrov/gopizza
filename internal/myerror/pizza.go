package myerror

import (
	"fmt"
)

var (
	ErrPizzasAlreadyExists = fmt.Errorf("pizza with this name already exists")
	ErrPizzaCreation       = fmt.Errorf("internal server error, failed to create a pizza")
	ErrPizzaNotFound       = fmt.Errorf("pizza not found")
	ErrPizzaFailedToLoad   = fmt.Errorf("pizza failed to load")
	ErrPizzaUpdating       = fmt.Errorf("internal server error, failed to update a pizza")
	ErrPizzaDeletion       = fmt.Errorf("internal server error, failed to delete a pizza")
)

package toasts

import "fmt"

var (
	ErrSaladsAlreadyExists = fmt.Errorf("salad with this name already exists")
	ErrSaladCreation       = fmt.Errorf("internal server error, failed to create a salad")
	ErrSaladNotFound       = fmt.Errorf("salad not found")
	ErrSaladFailedToLoad   = fmt.Errorf("salad failed to load")
	ErrSaladUpdating       = fmt.Errorf("internal server error, failed to update a salad")
	ErrSaladDeletion       = fmt.Errorf("internal server error, failed to delete a salad")
)

package toasts

import "fmt"

var (
	ErrCartItemCreation = fmt.Errorf("internal server error, failed to create a cart item")
	ErrCartDoesNotExist = fmt.Errorf("internal server error, cart does not exist")
)

package toasts

import "fmt"

var (
	ErrOrderCreation = fmt.Errorf("internal server error, failed to create a order")
	ErrOrderNotFound = fmt.Errorf("order not found")
)
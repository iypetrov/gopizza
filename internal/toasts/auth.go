package toasts

import "fmt"

var (
	ErrAuthEmailRequired    = fmt.Errorf("email is required")
	ErrAuthPasswordRequired = fmt.Errorf("password is required")
)
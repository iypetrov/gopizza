package toasts

import "fmt"

var (
	ErrUserAlreadyExists = fmt.Errorf("user with this name already exists")
	ErrUserCreation      = fmt.Errorf("internal server error, failed to create a user")
	ErrUserConfirmation  = fmt.Errorf("internal server error, failed to confirm a user")
	ErrUserNotFound      = fmt.Errorf("user not found")
)

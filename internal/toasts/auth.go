package toasts

import "fmt"

var (
	ErrAuthEmailRequired    = fmt.Errorf("email is required")
	ErrAuthPasswordRequired = fmt.Errorf("password is required")
	ErrAuthAddressRequired = fmt.Errorf("address is required")
	ErrAuthVerificationCodeNotCorrectFormat = fmt.Errorf("verification code is not in correct format")
)

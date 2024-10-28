package toasts

import (
	"fmt"
	"net/http"
)

var (
	ErrNameIsRequired              = fmt.Errorf("name is required")
	ErrImageIsRequired             = fmt.Errorf("image url is required")
	ErrPriceShouldBePositiveNumber = fmt.Errorf("price should be positive number")
	ErrNotValidUUID                = fmt.Errorf("not valid uuid")
	ErrImageNotFound               = fmt.Errorf("image not found")
	ErrDatabaseTransactionFailed   = fmt.Errorf("database transaction failed")
	ErrNotOwnAccount               = fmt.Errorf("not own account")
	ErrNotValidCookie              = fmt.Errorf("not valid cookie")
)

func ErrorFailedRender() Toast {
	return Toast{
		Message:    "failed to render component",
		StatusCode: http.StatusInternalServerError,
	}
}

func ErrorNotFound(err error) Toast {
	return Toast{
		Message:    err.Error(),
		StatusCode: http.StatusNotFound,
	}
}

func ErrorInternalServerError(err error) Toast {
	return Toast{
		Message:    err.Error(),
		StatusCode: http.StatusInternalServerError,
	}
}

func ErrorStatusUnauthorized(err error) Toast {
	return Toast{
		Message:    err.Error(),
		StatusCode: http.StatusUnauthorized,
	}
}

func ErrorStatusForbidden(err error) Toast {
	return Toast{
		Message:    err.Error(),
		StatusCode: http.StatusForbidden,
	}
}

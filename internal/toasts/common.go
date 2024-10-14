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
)

func ErrorFailedRender() Toast {
	return Toast{
		Message:    "failed to render component",
		StatusCode: http.StatusInternalServerError,
	}
}

func ErrorInternalServerError(err error) Toast {
	return Toast{
		Message:    err.Error(),
		StatusCode: http.StatusInternalServerError,
	}
}

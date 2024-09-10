package common

import (
	"fmt"
)

var (
	ErrNameIsRequired              = fmt.Errorf("name is required")
	ErrImageUrlIsRequired          = fmt.Errorf("image url is required")
	ErrPriceShouldBePositiveNumber = fmt.Errorf("price should be positive number")
)

func ErrorFailedRender() Toast {
	return Toast{
		Message:    fmt.Sprint("Failed to render component"),
		StatusCode: "warning",
	}
}

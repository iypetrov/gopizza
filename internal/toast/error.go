package toast

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrNameIsRequired              = fmt.Errorf("name is required")
	ErrImageUrlIsRequired          = fmt.Errorf("image url is required")
	ErrPriceShouldBePositiveNumber = fmt.Errorf("price should be positive number")
	ErrPizzasAlreadyExists         = fmt.Errorf("pizza with this name already exists")
	ErrPizzaCreation               = fmt.Errorf("internal server error, failed to create a pizza")
	ErrPizzaNotFound               = fmt.Errorf("pizza not found")
	ErrPizzaFailedToLoad           = fmt.Errorf("pizza failed to load")
	ErrPizzaUpdating               = fmt.Errorf("internal server error, failed to update a pizza")
	ErrPizzaDeletion               = fmt.Errorf("internal server error, failed to delete a pizza")
)

type CustomError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (e CustomError) Error() string {
	return fmt.Sprintf("custom error: %s", e.Message)
}

func ErrorInvalidRequestData(errs []error) CustomError {
	return CustomError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    strings.Join(errorsToStrings(errs), ", "),
	}
}

func ErrorBadRequest(err error) CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    err.Error(),
	}
}

func ErrorNotFound(err error) CustomError {
	return CustomError{
		StatusCode: http.StatusNotFound,
		Message:    err.Error(),
	}
}

func ErrorInternalServerError(err error) CustomError {
	return CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
	}
}

func ErrorInvalidJSON() CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprint("Invalid JSON request data"),
	}
}

func ErrorInvalidUUID() CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprint("Invalid UUID"),
	}
}

func ErrorFailedReadRequestBody() CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprint("Failed to read a request body"),
	}
}

func errorsToStrings(errors []error) []string {
	errs := make([]string, len(errors))
	for i, err := range errors {
		errs[i] = err.Error()
	}
	return errs
}

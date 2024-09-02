package utils

import (
	"fmt"
	"net/http"
	"strings"
)

type APIError struct {
	StatusCode int `json:"statusCode"`
	Message    any `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("API error: %d", e.StatusCode)
}

func InvalidRequestData(errors []error) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    strings.Join(errorsToStrings(errors), ", "),
	}
}

func BadRequest(err error) APIError {
	return APIError{
		StatusCode: http.StatusBadRequest,
		Message:    err.Error(),
	}
}

func NotFound(err error) APIError {
	return APIError{
		StatusCode: http.StatusNotFound,
		Message:    err.Error(),
	}
}

func InternalServerError(err error) APIError {
	return APIError{
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
	}
}

func InvalidJSON() APIError {
	return APIError{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprint("Invalid JSON request data"),
	}
}

func InvalidUUID() APIError {
	return APIError{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprint("Invalid UUID"),
	}
}

func FailedReadRequestBody() APIError {
	return APIError{
		StatusCode: http.StatusInternalServerError,
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
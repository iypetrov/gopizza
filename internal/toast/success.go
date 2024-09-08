package toast

import (
	"fmt"
	"net/http"
)

type CustomSuccess struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func SuccessPizzaCreated() CustomSuccess {
	return CustomSuccess{
		StatusCode: http.StatusCreated,
		Message:  fmt.Sprintf("pizza created successfully"),
	}
}

func SuccessPizzaUpdated() CustomSuccess {
	return CustomSuccess{
		StatusCode: http.StatusOK,
		Message:  fmt.Sprintf("pizza updated successfully"),
	}
}

func SuccessPizzaDeleted() CustomSuccess {
	return CustomSuccess{
		StatusCode: http.StatusOK,
		Message:  fmt.Sprintf("pizza deleted successfully"),
	}
}

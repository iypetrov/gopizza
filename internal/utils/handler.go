package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func MakeHandler(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			var apiErr APIError
			if errors.As(err, &apiErr) {
				err := WriteJson(w, apiErr.StatusCode, apiErr)
				if err != nil {
					return
				}
			}
		}
	}
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

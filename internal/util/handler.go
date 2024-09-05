package util

import (
	"encoding/json"
	"errors"
	"github.com/iypetrov/gopizza/internal/myerror"
	"net/http"
)

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func Make(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			var apiErr myerror.APIError
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

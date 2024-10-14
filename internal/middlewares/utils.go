package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/toasts"
)

type MiddlewareKey string

var UUIDKey MiddlewareKey = "UUID_KEY"

func UUIDFormat(next http.Handler) http.Handler {
	var val uuid.UUID
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if len(id) != 0 {
			i, err := uuid.Parse(id)
			if err != nil {
				err = writeErr(w, toasts.ErrorInternalServerError(err))
				if err != nil {
					return
				}
			}
			val = i
		}
		ctx := context.WithValue(r.Context(), UUIDKey, val)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func writeErr(w http.ResponseWriter, err toasts.Toast) error {
	w.WriteHeader(err.StatusCode)
	return json.NewEncoder(w).Encode(err)
}
package middleware

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/utils"
	"io"
	"net/http"
)

var (
	UUIDKey = "UUID_KEY"
	BodyKey = "BODY_KEY"
)

func UUIDFormat(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if len(id) != 0 {
			i, err := uuid.Parse(id)
			if err != nil {
				writeAPIError(w, utils.InvalidUUID())
				return
			}

			ctx := context.WithValue(r.Context(), UUIDKey, i)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}

func BodyFormat(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			writeAPIError(w, utils.FailedReadRequestBody())
			return
		}
		defer r.Body.Close()

		ctx := context.WithValue(r.Context(), BodyKey, body)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func writeAPIError(w http.ResponseWriter, apiErr utils.APIError) {
	w.WriteHeader(apiErr.StatusCode)
	json.NewEncoder(w).Encode(apiErr)
}

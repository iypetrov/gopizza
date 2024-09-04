package middleware

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/utils"
	"net/http"
)

func UUIDOrNotFound(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if len(id) != 0 {
			i, err := uuid.Parse(id)
			if err != nil {
				fmt.Println("error parsing uuid: ", err)
				utils.RedirectNotFoundView(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), UUIDKey, i)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}

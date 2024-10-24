package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/common"
	"github.com/iypetrov/gopizza/internal/toasts"
)

type MiddlewareKey string

var UUIDKey MiddlewareKey = "UUID_KEY"
var CookieName MiddlewareKey = MiddlewareKey(common.CookieName)

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

func AuthClient(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userCookie, err := common.ReadCookie(r)
		if err != nil {
			common.HxRedirect(w, "/login")
		}
		ctx := context.WithValue(r.Context(), CookieName, userCookie)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userCookie, err := common.ReadCookie(r)
		if err != nil {
			common.HxRedirect(w, "/login")
			return
		}
		if !userCookie.IsAdmin() {
			common.HxRedirect(w, "/login")
			return
		}
		ctx := context.WithValue(r.Context(), CookieName, userCookie)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func writeErr(w http.ResponseWriter, err toasts.Toast) error {
	w.WriteHeader(err.StatusCode)
	return json.NewEncoder(w).Encode(err)
}

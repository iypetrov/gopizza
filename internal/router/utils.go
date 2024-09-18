package router

import (
	"errors"
	"github.com/iypetrov/gopizza/internal/toasts"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func Make(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			var t toasts.Toast
			if errors.As(err, &t) {
				toasts.AddToast(w, t)
			}
		}
	}
}

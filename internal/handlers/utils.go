package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/iypetrov/gopizza/internal/toasts"
)

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	w.Header().Set("Content-Type", "text/html")

	err := c.Render(r.Context(), w)
	if err != nil {
		return toasts.ErrorFailedRender()
	}

	return nil
}

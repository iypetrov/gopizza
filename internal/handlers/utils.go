package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/dtos"
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

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func IsOwnAccount(userID uuid.UUID, cookie dtos.UserCookie) bool {
	return userID.String() == cookie.ID
}

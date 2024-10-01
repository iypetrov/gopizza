package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/iypetrov/gopizza/configs"
	"github.com/iypetrov/gopizza/internal/toasts"
)

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	err := c.Render(r.Context(), w)
	if err != nil {
		return toasts.ErrorFailedRender()
	}

	return nil
}

func hxRedirect(w http.ResponseWriter, path string) {
	w.Header().Set("HX-Redirect", fmt.Sprintf("%s%s", configs.Get().GetBaseWebUrl(), path))
}

func RedirectHomePage(w http.ResponseWriter) {
	hxRedirect(w, "/home")
}

func RedirectAdminHomePage(w http.ResponseWriter) {
	hxRedirect(w, fmt.Sprintf("%s%s", configs.Get().GetAdminPrefix(), "/home"))
}

package handlers

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/iypetrov/gopizza/configs"
	"github.com/iypetrov/gopizza/internal/toasts"
	"net/http"
)

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	err := c.Render(r.Context(), w)
	if err != nil {
		return toasts.ErrorFailedRender()
	}

	return nil
}

func hxRedirect(r *http.Request, path string) {
	r.Header.Set("HX-Redirect", fmt.Sprintf("%s%s", configs.Get().GetBaseWebUrl(), path))
}

func RedirectHomePage(r *http.Request) {
	hxRedirect(r, "/home")
}

func RedirectAdminHomePage(r *http.Request) {
	hxRedirect(r, fmt.Sprintf("%s%s", configs.Get().GetAdminPrefix(), "/home"))
}



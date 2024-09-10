package common

import (
	"github.com/a-h/templ"
	"net/http"
)

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) {
	err := c.Render(r.Context(), w)
	if err != nil {
		AddToast(w, ErrorFailedRender())
	}
}

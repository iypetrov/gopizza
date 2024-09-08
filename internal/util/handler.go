package util

import (
	"errors"
	"github.com/a-h/templ"
	"github.com/iypetrov/gopizza/internal/toast"
	"github.com/iypetrov/gopizza/web/template/component"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func Make(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			var terr toast.CustomError
			if errors.As(err, &terr) {
				RenderError(w, r, terr)
			}
		}
	}
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}

func RenderSuccess(w http.ResponseWriter, r *http.Request, suc toast.CustomSuccess) error {
	w.WriteHeader(suc.StatusCode)
	return Render(w, r, component.ToastSuccess(suc))
}

func RenderError(w http.ResponseWriter, r *http.Request, err toast.CustomError) error {
	w.WriteHeader(err.StatusCode)
	return Render(w, r, component.ToastError(err))
}

func RenderWarning(w http.ResponseWriter, r *http.Request, warn toast.CustomWarning) error {
	w.WriteHeader(warn.StatusCode)
	return Render(w, r, component.ToastWarning(warn))
}

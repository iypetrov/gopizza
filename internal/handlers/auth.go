package handlers

import (
	"github.com/iypetrov/gopizza/internal/dtos"
	"github.com/iypetrov/gopizza/internal/toasts"
	"github.com/iypetrov/gopizza/templates/components"

	"net/http"
)

type Auth struct{}

func NewAuth() Auth {
	return Auth{}
}

func (hnd *Auth) Login(w http.ResponseWriter, r *http.Request) error {
	req, err := dtos.ParseToLoginRequest(r)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.LoginForm(req, make(map[string]string)))
	}

	errs := req.Validate()
	if len(errs) > 0 {
		return Render(w, r, components.LoginForm(req, errs))
	}

	toasts.AddToast(w, toasts.Toast{
		Message:    "you are in",
		StatusCode: http.StatusCreated,
	})
	return Render(w, r, components.LoginForm(dtos.LoginRequest{}, make(map[string]string)))
}

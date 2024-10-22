package handlers

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/dtos"
	"github.com/iypetrov/gopizza/internal/services"
	"github.com/iypetrov/gopizza/internal/toasts"
	"github.com/iypetrov/gopizza/templates/components"

	"net/http"
)

type Auth struct {
	srv services.Auth
}

func NewAuth(srv services.Auth) Auth {
	return Auth{
		srv: srv,
	}
}

func (hnd *Auth) Register(w http.ResponseWriter, r *http.Request) error {
	req, err := dtos.ParseToRegisterRequest(r)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.RegisterForm(req, make(map[string]string)))
	}

	errs := req.Validate()
	if len(errs) > 0 {
		return Render(w, r, components.RegisterForm(req, errs))
	}

	id, err := hnd.srv.CreateUser(r.Context(), req.Email, req.Password, req.Address)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.RegisterForm(req, make(map[string]string)))
	}

	hxRedirect(w, fmt.Sprintf("/verification-code?id=%s&email=%s", id, req.Email))
	return Render(w, r, components.RegisterForm(dtos.RegisterRequest{}, make(map[string]string)))
}

func (hnd *Auth) VerifyRegistrationCode(w http.ResponseWriter, r *http.Request) error {
	req, err := dtos.ParseToRegisterVerificationRequest(r)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.RegisterVerificationForm(req, make(map[string]string)))
	}

	emptyReq := dtos.RegisterVerificationRequest{
		ID:    req.ID,
		Email: req.Email,
	}

	errs := req.Validate()
	if len(errs) > 0 {
		return Render(w, r, components.RegisterVerificationForm(emptyReq, errs))
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		return Render(w, r, components.RegisterVerificationForm(emptyReq, errs))
	}

	code := strings.Join([]string{
		req.CodeSymbol1,
		req.CodeSymbol2,
		req.CodeSymbol3,
		req.CodeSymbol4,
		req.CodeSymbol5,
		req.CodeSymbol6,
	},
		"",
	)

	err = hnd.srv.VerifyUserCode(
		r.Context(),
		id,
		req.Email,
		code,
	)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.RegisterVerificationForm(emptyReq, make(map[string]string)))
	}

	hxRedirect(w, "/login")
	return Render(w, r, components.RegisterVerificationForm(dtos.RegisterVerificationRequest{}, make(map[string]string)))
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

	hxRedirect(w, "/home")
	return Render(w, r, components.LoginForm(dtos.LoginRequest{}, make(map[string]string)))
}

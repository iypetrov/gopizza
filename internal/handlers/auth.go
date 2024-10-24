package handlers

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/common"
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

	common.HxRedirect(w, fmt.Sprintf("/verification-code?id=%s&email=%s", id, req.Email))
	return Render(w, r, components.RegisterForm(dtos.RegisterRequest{}, make(map[string]string)))
}

func (hnd *Auth) VerifyRegistrationCode(w http.ResponseWriter, r *http.Request) error {
	req, err := dtos.ParseToRegisterVerificationRequest(r)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.RegisterVerificationForm(*r, req, make(map[string]string)))
	}

	errs := req.Validate()
	if len(errs) > 0 {
		return Render(w, r, components.RegisterVerificationForm(*r, dtos.RegisterVerificationRequest{}, errs))
	}

	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		return Render(w, r, components.RegisterVerificationForm(*r, dtos.RegisterVerificationRequest{}, errs))
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
		r.URL.Query().Get("email"),
		code,
	)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.RegisterVerificationForm(*r, dtos.RegisterVerificationRequest{}, errs))
	}

	common.HxRedirect(w, "/login")
	return Render(w, r, components.RegisterVerificationForm(*r, dtos.RegisterVerificationRequest{}, make(map[string]string)))
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

	cookie, err := hnd.srv.VerifyUser(r.Context(), req.Email, req.Password)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.LoginForm(req, make(map[string]string)))
	}

	err = common.WriteCookie(w, cookie)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.LoginForm(req, make(map[string]string)))
	}

	if cookie.IsAdmin() {
		common.HxRedirect(w, "/admin/home")
	} else {
		common.HxRedirect(w, "/home")
	}

	return Render(w, r, components.LoginForm(dtos.LoginRequest{}, make(map[string]string)))
}

package dtos

import (
	"net/http"

	"github.com/iypetrov/gopizza/internal/toasts"
)

type LoginRequest struct {
	Email    string
	Password string
}

func (req *LoginRequest) Validate() map[string]string {
	errs := make(map[string]string)

	if len(req.Email) == 0 {
		errs["email"] = toasts.ErrAuthEmailRequired.Error()
	}

	if len(req.Password) == 0 {
		errs["password"] = toasts.ErrAuthPasswordRequired.Error()
	}

	return errs
}

func ParseToLoginRequest(r *http.Request) (LoginRequest, error) {
	err := r.ParseForm()
	if err != nil {
		return LoginRequest{}, err
	}

	var req LoginRequest
	req.Email = parseString(r, "email")
	req.Password = parseString(r, "password")

	return req, nil
}

type RegisterRequest struct {
	Email    string
	Password string
	Address	string
}

func (req RegisterRequest) Validate() map[string]string {
	errs := make(map[string]string)

	if len(req.Email) == 0 {
		errs["email"] = toasts.ErrAuthEmailRequired.Error()
	}

	if len(req.Password) == 0 {
		errs["password"] = toasts.ErrAuthPasswordRequired.Error()
	}

	if len(req.Address) == 0 {
		errs["address"] = toasts.ErrAuthAddressRequired.Error()
	}

	return errs
}

func ParseToRegisterRequest(r *http.Request) (RegisterRequest, error) {
	err := r.ParseForm()
	if err != nil {
		return RegisterRequest{}, err
	}

	var req RegisterRequest
	req.Email = parseString(r, "email")
	req.Password = parseString(r, "password")
	req.Address = parseString(r, "address")

	return req, nil
}

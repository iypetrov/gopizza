package dtos

import (
	"net/http"

	"github.com/iypetrov/gopizza/internal/toasts"
)

type RegisterRequest struct {
	Email    string
	Password string
	Address  string
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

type RegisterVerificationRequest struct {
	CodeSymbol1 string
	CodeSymbol2 string
	CodeSymbol3 string
	CodeSymbol4 string
	CodeSymbol5 string
	CodeSymbol6 string
}

func (req *RegisterVerificationRequest) Validate() map[string]string {
	errs := make(map[string]string)

	if len(req.CodeSymbol1) != 1 ||
		len(req.CodeSymbol2) != 1 ||
		len(req.CodeSymbol3) != 1 ||
		len(req.CodeSymbol4) != 1 ||
		len(req.CodeSymbol5) != 1 ||
		len(req.CodeSymbol6) != 1 {
		errs["code"] = toasts.ErrAuthVerificationCodeNotCorrectFormat.Error()
	}

	return errs
}

func ParseToRegisterVerificationRequest(r *http.Request) (RegisterVerificationRequest, error) {
	err := r.ParseForm()
	if err != nil {
		return RegisterVerificationRequest{}, err
	}

	var req RegisterVerificationRequest
	req.CodeSymbol1 = parseString(r, "codeSymbol1")
	req.CodeSymbol2 = parseString(r, "codeSymbol2")
	req.CodeSymbol3 = parseString(r, "codeSymbol3")
	req.CodeSymbol4 = parseString(r, "codeSymbol4")
	req.CodeSymbol5 = parseString(r, "codeSymbol5")
	req.CodeSymbol6 = parseString(r, "codeSymbol6")

	return req, nil
}

type UserCookie struct {
	ID           string
	Email        string
	AccessToken  string
	RefreshToken string
}

func (cookie *UserCookie) IsAdmin() bool {
	return cookie.Email == "ilia.yavorov.petrov@gmail.com"
}

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

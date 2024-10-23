package common

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"net/http"
	"strings"

	"github.com/iypetrov/gopizza/internal/dtos"
	"github.com/iypetrov/gopizza/internal/toasts"
)

var CookieName string = "GOPIZZA_COOKIE"

func WriteCookie(w http.ResponseWriter, value dtos.UserCookie) error {
	var buf bytes.Buffer

	err := gob.NewEncoder(&buf).Encode(&value)
	if err != nil {
		return err
	}

	cookie := http.Cookie{
		Name:     CookieName,
		Value:    buf.String(),
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	return write(w, cookie)
}

func ReadCookie(r *http.Request) (dtos.UserCookie, error) {
	value, err := read(r, CookieName)
	if err != nil {
		return dtos.UserCookie{}, err
	}

	reader := strings.NewReader(value)

	var userCookie dtos.UserCookie
	if err := gob.NewDecoder(reader).Decode(&userCookie); err != nil {
		return dtos.UserCookie{}, err
	}

	return userCookie, nil
}

func write(w http.ResponseWriter, cookie http.Cookie) error {
	cookie.Value = base64.URLEncoding.EncodeToString([]byte(cookie.Value))

	if len(cookie.String()) > 4096 {
		return toasts.ErrCookieValueTooLong
	}

	http.SetCookie(w, &cookie)

	return nil
}

func read(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}

	value, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		return "", toasts.ErrCookieInvalidValue
	}

	return string(value), nil
}

package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/iypetrov/gopizza/configs"
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

	return writeEncrypted(w, cookie, []byte(configs.Get().App.Secret))
}

func ReadCookie(r *http.Request) (dtos.UserCookie, error) {
	value, err := readEncrypted(r, CookieName, []byte(configs.Get().App.Secret))
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

func writeEncrypted(w http.ResponseWriter, cookie http.Cookie, secretKey []byte) error {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return err
	}

	plaintext := fmt.Sprintf("%s:%s", cookie.Name, cookie.Value)

	encryptedValue := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)

	cookie.Value = string(encryptedValue)

	return write(w, cookie)
}

func readEncrypted(r *http.Request, name string, secretKey []byte) (string, error) {
	encryptedValue, err := read(r, name)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()

	if len(encryptedValue) < nonceSize {
		return "", toasts.ErrCookieInvalidValue
	}

	nonce := encryptedValue[:nonceSize]
	ciphertext := encryptedValue[nonceSize:]

	plaintext, err := aesGCM.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		return "", toasts.ErrCookieInvalidValue
	}

	expectedName, value, ok := strings.Cut(string(plaintext), ":")
	if !ok {
		return "", toasts.ErrCookieInvalidValue
	}

	if expectedName != name {
		return "", toasts.ErrCookieInvalidValue
	}

	return value, nil
}

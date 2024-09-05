package util

import (
	"fmt"
	"github.com/iypetrov/gopizza/internal/config"
	"net/http"
)

func RedirectNotFoundView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/")
	http.Redirect(w, r, fmt.Sprintf("%s/404", config.Get().GetBaseWebUrl()), http.StatusTemporaryRedirect)
}

func RedirectHomeView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/")
	http.Redirect(w, r, fmt.Sprintf("%s/home", config.Get().GetBaseWebUrl()), http.StatusTemporaryRedirect)
}

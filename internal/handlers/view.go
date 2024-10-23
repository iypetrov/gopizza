package handlers

import (
	"net/http"

	"github.com/iypetrov/gopizza/templates/views"
)

func NotFoundView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.NotFound())
}

func RegisterView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.Register())
}

func RegisterVerificationView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.RegisterVerification(*r))
}

func LoginView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.Login())
}

func HomeView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.Home())
}

func AdminHomeView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.AdminHome())
}

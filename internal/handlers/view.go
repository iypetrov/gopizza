package handlers

import (
	"github.com/iypetrov/gopizza/templates/views"
	"net/http"
)

func NotFoundView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.NotFound())
}

func HomeView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.Home())
}

func AdminHomeView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.AdminHome())
}

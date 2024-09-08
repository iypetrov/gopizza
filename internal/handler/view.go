package handler

import (
	"github.com/iypetrov/gopizza/internal/util"
	"github.com/iypetrov/gopizza/web/template/view"
	"net/http"
)

func NotFoundView(w http.ResponseWriter, r *http.Request) {
	util.Render(w, r, view.NotFound())
}

func HomeView(w http.ResponseWriter, r *http.Request) {
	util.Render(w, r, view.Home())
}

func AdminHomeView(w http.ResponseWriter, r *http.Request) {
	util.Render(w, r, view.AdminHome())
}

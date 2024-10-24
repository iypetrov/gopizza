package common

import (
	"fmt"
	"net/http"

	"github.com/iypetrov/gopizza/configs"
)

func HxRedirect(w http.ResponseWriter, path string) {
	w.Header().Set("HX-Redirect", fmt.Sprintf("%s%s", configs.Get().GetBaseWebUrl(), path))
}

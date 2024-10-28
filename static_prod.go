//go:build prod
// +build prod

package main

import (
	"embed"
	"net/http"
)

//go:embed public
var publicFS embed.FS

func Public() http.Handler {
	return http.FileServerFS(publicFS)
}

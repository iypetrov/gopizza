package dtos

import (
	"net/http"
	"strconv"
)

func parseString(r *http.Request, key string) string {
	return r.FormValue(key)
}

func parseBool(r *http.Request, key string) bool {
	val, err := strconv.ParseBool(r.FormValue(key))
	if err != nil {
		val = false
	}
	return val
}

func parseFloat(r *http.Request, key string) float64 {
	val, err := strconv.ParseFloat(r.FormValue(key), 64)
	if err != nil {
		val = 0
	}
	return val
}

package lib

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Slugify(s string) string {
	return s
}

func StringURLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}
func IntURLParam(r *http.Request, key string) (int, error) {
	param := chi.URLParam(r, key)
	return strconv.Atoi(param)
}

package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func intURlParam(r *http.Request, key string) (int, error) {
	s := chi.URLParam(r, key)
	return strconv.Atoi(s)
}

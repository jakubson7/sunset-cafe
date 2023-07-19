package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func IntURlParam(r *http.Request, key string) (int, error) {
	s := chi.URLParam(r, key)
	return strconv.Atoi(s)
}

func Format(s string, args ...any) []byte {
	return []byte(fmt.Sprintf(s, args...))
}

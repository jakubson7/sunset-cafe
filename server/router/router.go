package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Group(testGroup)

	return r
}

func Serve(r *chi.Mux) {
	http.ListenAndServe(":3000", r)
}

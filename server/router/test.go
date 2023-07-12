package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jakubson7/sunset-cafe/repos"
)

func testGroup(r chi.Router) {
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		meals, err := repos.Meal.GetAll()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(meals)

		w.Write([]byte("Hello world"))
	})
}

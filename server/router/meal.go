package router

import (
	"fmt"
	"net/http"

	"github.com/jakubson7/sunset-cafe/lib"
	"github.com/jakubson7/sunset-cafe/repos"
)

func handleGetMealByID(w http.ResponseWriter, r *http.Request) {
	ID, err := lib.IntURLParam(r, "ID")

	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte("{ID} is not a int"))
		return
	}

	meal, err := repos.Meal.GetByID(ID)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Meal can't be found"))
	}

	w.Write([]byte(fmt.Sprintf("%v", meal)))
}

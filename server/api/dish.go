package api

import (
	"encoding/json"
	"net/http"

	"github.com/jakubson7/sunset-cafe/models"
)

func (api *API) handleCreateDish(w http.ResponseWriter, r *http.Request) {
	err := api.dish.CreateOne(models.NewDish("Spaghetti", 20, "Pretty italian food i have to say.", 1))
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		w.Write(Format("%v", err))
	}
}

func (api *API) handleGetDishByID(w http.ResponseWriter, r *http.Request) {
	ID, err := IntURlParam(r, "ID")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Format("%v", err))
	}

	dish, err := api.dish.GetByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		w.Write(Format("%v", err))
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dish)
}

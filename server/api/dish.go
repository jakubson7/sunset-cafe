package api

import "net/http"

func (api *Api) handleGetDishByID(w http.ResponseWriter, r *http.Request) {
	ID, err := intURlParam(r, "ID")
	if err != nil {
		sendErr(w, err)
		return
	}

	dish, err := api.dishService.GetDishByID(int64(ID))
	if err != nil {
		sendErr(w, err)
		return
	}

	sendData(w, dish)
	return
}

func (api *Api) handleGetDishes(w http.ResponseWriter, r *http.Request) {
	dishes, err := api.dishService.GetDishes(10, 0)
	send(w, dishes, err)
	return
}

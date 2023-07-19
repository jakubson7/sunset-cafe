package api

import (
	"net/http"

	"github.com/jakubson7/sunset-cafe/models"
)

func (api *API) handleCreateImage(w http.ResponseWriter, r *http.Request) {
	err := api.image.CreateOne(models.NewImage("Spaghetti", "LOCAL", "", "", ""))
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		w.Write(Format("%v", err))
	}
}

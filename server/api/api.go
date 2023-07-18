package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jakubson7/sunset-cafe/db"
	"github.com/jakubson7/sunset-cafe/models"
)

type API struct {
	addr  string
	db    *sql.DB
	dish  *models.DishModel
	image *models.ImageModel
}

func NewAPI() *API {
	db := db.NewMemorySqlite()
	dishModel := models.NewDishModel(db)
	imageModel := models.NewImageModel(db)

	api := &API{
		addr:  ":8000",
		db:    db,
		dish:  dishModel,
		image: imageModel,
	}

	return api
}

func (api *API) Start() {
	router := chi.NewRouter()

	err := http.ListenAndServe(api.addr, router)
	if err != nil {
		log.Fatal(err)
	}
}

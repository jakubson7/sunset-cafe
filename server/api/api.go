package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jakubson7/sunset-cafe/db"
	"github.com/jakubson7/sunset-cafe/models"
	"github.com/jakubson7/sunset-cafe/storage"
)

type API struct {
	addr    string
	db      *sql.DB
	storage storage.StorageProvider
	dish    *models.DishModel
	image   *models.ImageModel
}

func NewAPI() *API {
	api := &API{}

	api.addr = ":8000"
	api.db = db.NewMemorySqlite()
	api.storage = storage.NewLocalStorage("/storage", api.addr)
	api.dish = models.NewDishModel(api.db)
	api.image = models.NewImageModel(api.db)

	return api
}

func (api *API) Start() {
	api.dish.SetupTable()
	api.image.SetupTable()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/dish", api.handleCreateDish)
	router.Get("/dish/{ID}", api.handleGetDishByID)
	router.Get("/image", api.handleCreateImage)

	err := http.ListenAndServe(api.addr, router)
	if err != nil {
		log.Fatal(err)
	}
}

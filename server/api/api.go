package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jakubson7/sunset-cafe/models"
	"github.com/jakubson7/sunset-cafe/services"
)

type Api struct {
	addr           string
	logger         *log.Logger
	sqliteService  *services.SqliteService
	userService    *services.UserService
	dishService    *services.DishService
	productService *services.ProductService
}

func NewApi() *Api {
	api := new(Api)

	api.addr = ":8080"
	api.logger = log.Default()
	api.sqliteService = services.NewMemorySqliteService()
	api.sqliteService.CreateTables()

	api.userService = services.NewUserService(api.sqliteService)
	api.dishService = services.NewDishService(api.sqliteService)
	api.productService = services.NewProductService(api.sqliteService)

	return api
}

func (api *Api) Start() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", api.handleGetRoot)
	r.Get("/dish/{ID}", api.handleGetDishByID)
	r.Get("/dish", api.handleGetDishes)

	if err := http.ListenAndServe(api.addr, r); err != nil {
		api.logger.Fatal(err)
	}

	api.sqliteService.Close()
}

func (api *Api) Mock() {
	var err error
	_, err = api.productService.CreateProduct(models.ProductParams{"Flower"})
	_, err = api.productService.CreateProduct(models.ProductParams{"Tomato"})
	_, err = api.productService.CreateProduct(models.ProductParams{"Water"})
	_, err = api.productService.CreateProduct(models.ProductParams{"Olive oil"})
	_, err = api.dishService.CreateDish(models.DishParams{
		Name:        "Pizza",
		Description: "Tasty nad traditional italian dish",
		Price:       35,
		Images:      []models.DishImageParams{},
		Ingredients: []models.IngredientParams{
			models.IngredientParams{1},
			models.IngredientParams{2},
			models.IngredientParams{3},
			models.IngredientParams{4},
		},
	})
	_, err = api.dishService.CreateDish(models.DishParams{
		Name:        "Spaghetti",
		Description: "Tasty nad traditional italian dish",
		Price:       25,
		Images:      []models.DishImageParams{},
		Ingredients: []models.IngredientParams{
			models.IngredientParams{2},
			models.IngredientParams{3},
			models.IngredientParams{4},
		},
	})
	_, err = api.dishService.CreateDish(models.DishParams{
		Name:        "Water",
		Description: "Just water",
		Price:       5,
		Images:      []models.DishImageParams{},
		Ingredients: []models.IngredientParams{
			models.IngredientParams{3},
		},
	})
	if err != nil {
		api.logger.Fatal(err)
	}
}

func (api *Api) handleGetRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello! It's Sunset Cafe!"))
}

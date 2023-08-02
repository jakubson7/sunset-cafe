package main

import (
	"fmt"
	"log"

	"github.com/jakubson7/sunset-cafe/models"
	"github.com/jakubson7/sunset-cafe/services"
)

func main() {
	sqliteService := services.NewMemorySqliteService()
	sqliteService.CreateTables()

	dishService := services.NewDishService(sqliteService)
	productService := services.NewProductService(sqliteService)

	var err error
	_, err = productService.CreateProduct(models.ProductParams{"Flower"})
	_, err = productService.CreateProduct(models.ProductParams{"Tomato"})
	_, err = productService.CreateProduct(models.ProductParams{"Water"})
	_, err = productService.CreateProduct(models.ProductParams{"Olive oil"})
	if err != nil {
		log.Fatal(err)
	}

	products, err := productService.GetAllProducts()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", products)

	dish, err := dishService.CreateDish(models.DishParams{
		Name:        "Spaghetti",
		Description: "Tasty nad traditional italian dish",
		Price:       20,
		Images:      []models.DishImageParams{},
		Ingredients: []models.IngredientParams{
			models.IngredientParams{2},
			models.IngredientParams{3},
			models.IngredientParams{4},
		},
	})
	_, err = dishService.CreateDish(models.DishParams{
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
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", dish)

	dishes, err := dishService.GetDishes(10, 0)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(dishes)
}

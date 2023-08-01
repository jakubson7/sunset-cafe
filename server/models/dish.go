package models

import (
	"errors"
)

type DishParams struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Images      []DishImageParams  `json:"dishImages"`
	Ingredients []IngredientParams `json:"ingredients"`
}

type Dish struct {
	DishID      int64     `json:"dishID"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Images      []Image   `json:"images"`
	Ingredients []Product `json:"products"`
}

type DishImage struct {
	DishID  int64 `json:"dishID"`
	ImageID int64 `json:"imageID"`
}

type DishImageParams struct {
	ImageID int64 `json:"imageID"`
}

type Ingredient struct {
	DishID    int64 `json:"dishID"`
	ProductID int64 `json:"productID"`
}

type IngredientParams struct {
	ProductID int64 `json:"productID"`
}

const DishSQL = `
	CREATE TABLE dishes (
		dishID INTEGER,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		price REAL NOT NULL,

		PRIMARY KEY (dishID)
	)
`
const DishImageSQL = `
	CREATE TABLE dishImages (
		dishID INTEGER,
		imageID INTEGER,

		PRIMARY KEY (dishID, imageID),
		FOREIGN KEY (dishID) REFERENCES dishes(dishID),
		FOREIGN KEY (imageID) REFERENCES images(imageID)
	)
`
const IngredientSQL = `
	CREATE TABLE ingredients (
		dishID INTEGER,
		productID INTEGER,

		PRIMARY KEY (dishID, productID),
		FOREIGN KEY (dishID) REFERENCES dishes(dishID),
		FOREIGN KEY (productID) REFERENCES products(productID)
	)
`

func (params *DishParams) Validate() error {
	if isEmpty(params.Name) {
		return errors.New("name cannot be an empty string")
	}
	if params.Price <= 0 {
		return errors.New("price cannot be equal or smaller than 0")
	}

	return nil
}

package models

import (
	"errors"
)

type DishParams struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Dish struct {
	DishID int64 `json:"dishID"`
	DishParams
	Images      []Image   `json:"images"`
	Ingredients []Product `json:"ingredients"`
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
		FOREIGN KEY (imageID) REFERENCES images(imageID)
	)
`

func (create *DishParams) Validate() error {
	if isEmpty(create.Name) {
		return errors.New("name cannot be an empty string")
	}
	if create.Price <= 0 {
		return errors.New("price cannot be equal or smaller than 0")
	}

	return nil
}

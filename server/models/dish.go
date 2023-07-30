package models

import (
	"errors"
	"fmt"
)

type DishCreate struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Dish struct {
	DishID int64 `json:"dishID"`
	Timestamp
	DishCreate

	Images      []Image   `json:"images"`
	Ingredients []Product `json:"ingredients"`
}

const DishSQL = `
	CREATE TABLE dishes (
		dishID INTEGER,
		createdAt INTEGER NOT NULL,
		updatedAt INTEGER NOT NULL,
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

func (create *DishCreate) Validate() error {
	if isEmpty(create.Name) {
		return create.efrom("name cannot be an empty string")
	}
	if create.Price <= 0 {
		return create.efrom("price cannot be equal or smaller than 0")
	}

	return nil
}

func (create *DishCreate) efrom(text string) error {
	return errors.New(fmt.Sprintf("(DishCreate) -> %s", text))
}
func (create *Dish) efrom(text string) error {
	return errors.New(fmt.Sprintf("(Dish) -> %s", text))
}
func (create *Dish) ewrap(err error) error {
	if err == nil {
		return nil
	}
	return errors.New(fmt.Sprintf("(Dish) -> %v", err))
}

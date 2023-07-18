package models

import (
	"database/sql"

	"github.com/jakubson7/sunset-cafe/db"
)

type Dish struct {
	ID          int     `json:"ID"`
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	ImgID       int     `json:"imgID"`
}

type DishWithImage struct {
	Dish
	Image Image `json:"image"`
}

type DishModel struct {
	db *sql.DB
}

func NewDishModel(db *sql.DB) *DishModel {
	return &DishModel{db}
}

func (m *DishModel) SetupTable() error {
	return db.PrepareAndExec(m.db, `
		CREATE TABLE dishes (
			dishID INTEGER,
			name TEXT NOT NULL,
			slug TEXT NOT NULL,
			price FLOAT NOT NULL,
			desciption TEXT NOT NULL,
			imgID int,

			PRIMARY KEY(dishID),
			FOREIGN KEY(imgID) REFERENCES images(imgID)
		)
	`)
}

func (m *DishModel) GetByID(ID string) (*DishWithImage, error) {
	row, err := db.PrepareAndQueryOne(m.db, `
		SELECT *
		FROM dishes
		INNER JOIN images
		ON dishID = imageID
	`, ID)

	if err != nil {
		return nil, err
	}

	var dish DishWithImage
	err = row.Scan(
		&dish.ID, &dish.Name, &dish.Slug, &dish.Price, &dish.Description, &dish.ImgID,
		&dish.Image.ID, &dish.Image.Name, &dish.Image.Slug, &dish.Image.Provider, &dish.Image.SmallURL, &dish.Image.MediumURL, &dish.Image.BigURL,
	)
	if err != nil {
		return nil, err
	}

	return &dish, nil
}

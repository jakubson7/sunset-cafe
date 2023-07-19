package models

import (
	"database/sql"
	"log"

	"github.com/gosimple/slug"
	"github.com/jakubson7/sunset-cafe/db"
)

type Dish struct {
	ID          int     `json:"ID"`
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImgID       int     `json:"imgID"`
}

type DishWithImage struct {
	Dish
	Image Image `json:"image"`
}

type DishModel struct {
	db *sql.DB
}

func NewDish(name string, price float32, description string, imgID int) *Dish {
	return &Dish{
		Name:        name,
		Slug:        slug.Make(name),
		Price:       price,
		Description: description,
		ImgID:       imgID,
	}
}

func NewDishModel(db *sql.DB) *DishModel {
	return &DishModel{db}
}

func (m *DishModel) SetupTable() {
	err := db.PrepareAndExec(m.db, `
		CREATE TABLE dishes (
			dishID INTEGER,
			name TEXT NOT NULL,
			slug TEXT NOT NULL,
			desciption TEXT NOT NULL,
			price FLOAT NOT NULL,
			imgID INTEGER,

			PRIMARY KEY(dishID),
			FOREIGN KEY(imgID) REFERENCES images(imgID)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *DishModel) GetByID(ID int) (*DishWithImage, error) {
	row, err := db.PrepareAndQueryOne(m.db, `
		SELECT *
		FROM dishes
		INNER JOIN images
		ON dishes.dishID = images.imageID
		WHERE dishes.dishID = $1
	`, ID)

	if err != nil {
		return nil, err
	}

	var dish DishWithImage
	err = row.Scan(
		&dish.ID, &dish.Name, &dish.Slug, &dish.Description, &dish.Price, &dish.ImgID,
		&dish.Image.ID, &dish.Image.Name, &dish.Image.Slug, &dish.Image.Provider, &dish.Image.SmallURL, &dish.Image.MediumURL, &dish.Image.BigURL,
	)
	if err != nil {
		return nil, err
	}

	return &dish, nil
}

func (m *DishModel) CreateOne(dish *Dish) error {
	return db.PrepareAndExec(m.db, `
		INSERT INTO dishes
			(name, slug, desciption, price, imgID)
		VALUES
			($1, $2, $3, $4, $5)
	`, dish.Name, dish.Slug, dish.Description, dish.Price, dish.ImgID)
}

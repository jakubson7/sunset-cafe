package services

import (
	"database/sql"
	"log"
)

type DishService struct {
	db                *sql.DB
	createDish        *sql.Stmt
	getDishByID       *sql.Stmt
	getDishImageByID  *sql.Stmt
	getIngredientByID *sql.Stmt
	getDishImages     *sql.Stmt
	getIngredients    *sql.Stmt
}

func NewDishService(sqliteService *SqliteService) *DishService {
	s := &DishService{}
	var err error

	s.db = sqliteService.DB
	s.createDish, err = s.db.Prepare(`INSERT INTO dishes (name, price) VALUES ($1, $2)`)
	s.getDishByID, err = s.db.Prepare(`SELECT * FROM dishes WHERE dishID = $1`)
	s.getDishImageByID, err = s.db.Prepare(`
		SELECT images.* FROM dishes
			INNER JOIN dishImages ON dishImages.dishID = dishes.dishID
			INNER JOIN images ON images.imageID = dishImages.imageID
			WHERE dishes.dishID = $1
	`)
	s.getIngredientByID, err = s.db.Prepare(`
		SELECT products.* FROM dishes
			INNER JOIN ingredients ON ingredients.dishID = dishes.dishID
			INNER JOIN products ON products.productID = ingredients.productID
			WHERE dishes.dishID = $1
	`)
	s.getDishImages, err = s.db.Prepare(`
		SELECT query_dishID, images.* FROM dishes
			INNER JOIN dishImages ON dishImages.dishID = dishes.dishID
			INNER JOIN images ON images.imageID = dishImages.imageID
			WHERE query_dishID IN (
				SELECT dishes.dishID FROM dishes LIMIT $1 OFFSET $2
			)
	`)
	s.getIngredients, err = s.db.Prepare(`
		SELECT query_dishID, products.* FROM dishes
			INNER JOIN ingredients ON ingredients.dishID = dishes.dishID
			INNER JOIN products ON products.productID = ingredients.productID
			WHERE query_dishID IN (
				SELECT dishes.dishID FROM dishes LIMIT $1 OFFSET $2
			)
	`)

	if err != nil {
		log.Fatal(err)
	}

	return s
}

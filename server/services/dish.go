package services

import (
	"database/sql"
	"log"
)

type DishService struct {
	name              string
	db                *sql.DB
	createDish        *sql.Stmt
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
	s.getDishImageByID, err = s.db.Prepare(`SELECT * FROM dishes WHERE dishID = $1`)
	s.getIngredientByID, err = s.db.Prepare(``)
	s.getDishImages, err = s.db.Prepare(`SELECT * FROM dishes LIMIT $1 OFFSET $2`)
	s.getIngredients, err = s.db.Prepare(``)

	if err != nil {
		log.Fatal(err)
	}

	return s
}

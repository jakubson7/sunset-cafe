package services

import (
	"database/sql"
	"log"
)

type DishService struct {
	name        string
	db          *sql.DB
	createDish  *sql.Stmt
	getDishByID *sql.Stmt
	getDishes   *sql.Stmt
}

func NewDishService(sqliteService *SqliteService) *DishService {
	s := &DishService{}
	var err error

	s.db = sqliteService.DB
	s.createDish, err = s.db.Prepare(`INSERT INTO dishes (createdAt, updatedAt, name) VALUES ($1, $2, $3)`)
	s.getDishByID, err = s.db.Prepare(`SELECT * FROM dishes WHERE dishID = $1`)
	s.getDishes, err = s.db.Prepare(`SELECT * FROM dishes LIMIT $1 OFFSET $2`)

	if err != nil {
		log.Fatal(err)
	}

	return s
}

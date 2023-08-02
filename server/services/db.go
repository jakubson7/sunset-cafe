package services

import (
	"database/sql"
	"log"

	"github.com/jakubson7/sunset-cafe/models"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteService struct {
	DB *sql.DB
}

func NewMemorySqliteService() *SqliteService {
	s := &SqliteService{}

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	s.DB = db

	return s
}

func (s *SqliteService) CreateTables() {
	var err error

	_, err = s.DB.Exec(models.UserSQL)
	_, err = s.DB.Exec(models.DishSQL)
	_, err = s.DB.Exec(models.ImageSQL)
	_, err = s.DB.Exec(models.ProductSQL)
	_, err = s.DB.Exec(models.DishImageSQL)
	_, err = s.DB.Exec(models.IngredientSQL)

	if err != nil {
		log.Fatal(err)
	}
}

func (s *SqliteService) Close() {
	s.Close()
}

package services

import (
	"database/sql"
	"errors"
	"fmt"
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
		s.efatal(err)
	}

	s.DB = db

	return s
}

func (s *SqliteService) CreateTables() {
	_, err := s.DB.Exec(models.UserSQL)
	_, err = s.DB.Exec(models.DishSQL)

	if err != nil {
		s.efatal(err)
	}
}

func (s *SqliteService) ewrap(err error) error {
	return errors.New(fmt.Sprintf("(SqliteService) -> %v", err))
}
func (s *SqliteService) efatal(err error) {
	log.Fatal(s.ewrap(err))
}

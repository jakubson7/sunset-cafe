package services

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewMemorySqliteService() (*sql.DB, error) {
	e := newServiceError("MemorySqliteService", "NewMemorySqliteService")

	db, err := sql.Open("sqlite3", ":memory:")
	return db, e.Wrap(err)
}

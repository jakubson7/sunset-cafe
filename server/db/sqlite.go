package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	DB *sql.DB
}

func (sqlite *Sqlite) GetVersion() (string, error) {
	var version string
	err := sqlite.DB.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		return "", err
	}

	return version, nil
}

func (sqlite *Sqlite) Close() {
	sqlite.DB.Close()
}

func (Sqlite *Sqlite) QueryRaw(cb func(db *sql.DB) error) error {
	return cb(Sqlite.DB)
}

func CreateSqliteProvider() DBProvider {
	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatal(err)
	}

	sqlite := &Sqlite{
		DB: db,
	}

	return sqlite
}

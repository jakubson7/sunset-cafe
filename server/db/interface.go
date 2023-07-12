package db

import (
	"database/sql"

	"github.com/jakubson7/sunset-cafe/lib"
)

type DBProvider interface {
	Close()
	GetVersion() (string, error)
	SetupMeal() error
	CreateMeal(meal *lib.Meal) error
	QueryRaw(func(db *sql.DB) error) error
}

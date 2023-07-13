package repos

import (
	"database/sql"
	"log"
)

func Setup(db *sql.DB) {
	Meal.Setup(db)
}

func Refresh() {
	err := Meal.Refresh()

	if err != nil {
		log.Fatal(err)
	}
}

func prepareAndQuery(db *sql.DB, query string, args ...any) (*sql.Rows, error) {
	stm, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stm.Query(args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func prepareAndExec(db *sql.DB, query string, args ...any) error {
	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stm.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}

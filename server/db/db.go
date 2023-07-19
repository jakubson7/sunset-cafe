package db

import (
	"database/sql"
)

func PrepareAndExec(db *sql.DB, query string, args ...any) error {
	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}

	// Empty args are read as []inteface{} and couse error while beeing passed to `Exec`
	if len(args) == 0 {
		_, err := stm.Exec()
		return err
	}
	_, err = stm.Exec(args...)

	return err
}

func PrepareAndQuery(db *sql.DB, query string, args ...any) (*sql.Rows, error) {
	stm, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	return stm.Query(args...)
}

func PrepareAndQueryOne(db *sql.DB, query string, args ...any) (*sql.Row, error) {
	stm, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	return stm.QueryRow(args...), nil
}

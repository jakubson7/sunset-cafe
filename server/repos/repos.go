package repos

import "database/sql"

func Setup(db *sql.DB) {
	Meal.Setup(db)
}

func SetupFromScratch(db *sql.DB) {
	Meal.SetupFromScratch(db)
}

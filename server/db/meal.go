package db

import "github.com/jakubson7/sunset-cafe/lib"

func (sqlite *Sqlite) SetupMeal() error {
	sql := `
		--sql
		DROP TABLE IF EXISTS meals;
		CREATE TABLE meals (
  		MealID INTEGER PRIMARY KEY,
			Name TEXT,
			Slug TEXT,
			Price INTEGER,
			ImgID INTEGER
		);
	`

	_, err := sqlite.DB.Exec(sql)
	return err
}

func (sqlite *Sqlite) CreateMeal(meal *lib.Meal) error {

	return nil
}

package app

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jakubson7/sunset-cafe/lib"
)

func TestPlugin(app *App) {
	version, err := app.DBProvider.GetVersion()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(version)

	err = app.DBProvider.SetupMeal()
	if err != nil {
		log.Fatal(err)
	}

	app.DBProvider.QueryRaw(func(db *sql.DB) error {
		_, err := db.Exec(`
			--sql
			INSERT INTO meals (Name, Slug, Price, ImgID) VALUES ('Spaghetti', 'spaghetti', 10, 1);
		`)
		return err
	})
	app.DBProvider.QueryRaw(func(db *sql.DB) error {
		stm, err := db.Prepare(`
			--sql
			SELECT * FROM meals;
		`)

		if err != nil {
			return err
		}

		var meal lib.Meal
		err = stm.QueryRow().Scan(&meal.ID, &meal.Name, &meal.Slug, &meal.Price, &meal.ImgID)
		fmt.Print(meal)
		return err
	})
}

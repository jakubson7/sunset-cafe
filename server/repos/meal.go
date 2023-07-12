package repos

import (
	"database/sql"

	"github.com/jakubson7/sunset-cafe/lib"
)

type MealRepo struct {
	DB *sql.DB
}

func (repo *MealRepo) Setup(db *sql.DB) {
	repo.DB = db
}

func (repo *MealRepo) SetupFromScratch(db *sql.DB) error {
	repo.DB = db

	_, err := repo.DB.Exec(`
		--sql
		DROP TABLE IF EXISTS meals;
		CREATE TABLE meals (
			MealID INTEGER PRIMARY KEY,
			Name TEXT,
			Slug TEXT,
			Price INTEGER,
			ImgID INTEGER
		);
	`)
	return err
}

func scanMealSQLRows(rows *sql.Rows) ([]lib.Meal, error) {
	var meals []lib.Meal
	for rows.Next() {
		m := lib.Meal{}

		err := rows.Scan(&m.ID, &m.Name, &m.Slug, &m.Price, &m.ImgID)
		if err != nil {
			return nil, err
		}

		meals = append(meals, m)
	}
	return meals, nil
}

func (repo *MealRepo) GetAll() ([]lib.Meal, error) {
	stm, err := repo.DB.Prepare(`SELECT * FROM meals;`)
	if err != nil {
		return nil, err
	}

	rows, err := stm.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	meals, err := scanMealSQLRows(rows)
	if err != nil {
		return nil, err
	}

	return meals, nil
}

func (repo *MealRepo) GetByID(ID int) (*lib.Meal, error) {
	stm, err := repo.DB.Prepare(`SELECT * FROM meals WHERE MealID = $1;`)
	if err != nil {
		return nil, err
	}

	row, err := stm.Query(ID)
	if err != nil {
		return nil, err
	}

	meals, err := scanMealSQLRows(row)
	return &meals[0], err
}

func (repo *MealRepo) CreateOne(meal *lib.Meal) error {
	stm, err := repo.DB.Prepare(`INSERT INTO meals (Name, Slug, Price, ImgID) VALUES ($1, $2, $3, $4);`)
	if err != nil {
		return err
	}

	_, err = stm.Exec(meal.Name, meal.Slug, meal.Price, meal.ImgID)
	if err != nil {
		return err
	}

	return nil
}

var Meal MealRepo

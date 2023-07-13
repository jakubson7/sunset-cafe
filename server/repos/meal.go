package repos

import (
	"database/sql"

	"github.com/jakubson7/sunset-cafe/lib"
)

type MealRepo struct {
	DB *sql.DB
}

var Meal MealRepo

func (repo *MealRepo) Setup(db *sql.DB) {
	repo.DB = db
}

func (repo *MealRepo) Refresh() error {
	_, err := repo.DB.Exec(`
		--sql
		DROP TABLE IF EXISTS meals;
		CREATE TABLE meals (
			MealID INTEGER PRIMARY KEY,
			Name TEXT NOT NULL,
			Slug TEXT NOT NULL,
			Price INTEGER NOT NULL,
			ImgID INTEGER
		);
	`)
	return err
}

func extractMeals(rows *sql.Rows) ([]lib.Meal, error) {
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
	rows, err := prepareAndQuery(repo.DB, `SELECT * FROM meals;`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return extractMeals(rows)
}

func (repo *MealRepo) GetByID(ID int) (*lib.Meal, error) {
	rows, err := prepareAndQuery(repo.DB, `SELECT * FROM meals WHERE MealID = $1;`, ID)
	if err != nil {
		return nil, err
	}

	meals, err := extractMeals(rows)
	return &meals[0], err
}

func (repo *MealRepo) CreateOne(meal *lib.Meal) error {
	return prepareAndExec(
		repo.DB,
		`INSERT INTO meals (Name, Slug, Price, ImgID) VALUES ($1, $2, $3, $4);`,
		meal.Name, meal.Slug, meal.Price, meal.ImgID,
	)
}

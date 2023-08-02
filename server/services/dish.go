package services

import (
	"database/sql"
	"log"

	"github.com/jakubson7/sunset-cafe/models"
)

type DishService struct {
	db                  *sql.DB
	createDish          *sql.Stmt
	getDishByID         *sql.Stmt
	getDishImageByID    *sql.Stmt
	getIngredientByID   *sql.Stmt
	getDishes           *sql.Stmt
	getDishImages       *sql.Stmt
	getIngredients      *sql.Stmt
	sqlCreateDishImage  string
	sqlCreateIngredient string
}

func NewDishService(sqliteService *SqliteService) *DishService {
	s := &DishService{}
	var err error

	s.db = sqliteService.DB
	s.createDish, err = s.db.Prepare(`INSERT INTO dishes (name, description, price) VALUES ($1, $2, $3)`)
	s.sqlCreateDishImage = `INSERT INTO dishImages (dishID, imageID) VALUES ($1, $2)`
	s.sqlCreateIngredient = `INSERT INTO ingredients (dishID, productID) VALUES ($1, $2)`
	s.getDishByID, err = s.db.Prepare(`SELECT * FROM dishes WHERE dishID = $1`)
	s.getDishImageByID, err = s.db.Prepare(`
		SELECT images.* FROM dishes
			INNER JOIN dishImages ON dishImages.dishID = dishes.dishID
			INNER JOIN images ON images.imageID = dishImages.imageID
			WHERE dishes.dishID = $1
	`)
	s.getIngredientByID, err = s.db.Prepare(`
		SELECT products.* FROM dishes
			INNER JOIN ingredients ON ingredients.dishID = dishes.dishID
			INNER JOIN products ON products.productID = ingredients.productID
			WHERE dishes.dishID = $1
	`)
	s.getDishes, err = s.db.Prepare(`SELECT * FROM dishes LIMIT $1 OFFSET $2`)
	s.getDishImages, err = s.db.Prepare(`
		SELECT dishImages.dishID, images.* FROM dishes
			INNER JOIN dishImages ON dishImages.dishID = dishes.dishID
			INNER JOIN images ON images.imageID = dishImages.imageID
			WHERE dishes.dishID IN (
				SELECT dishes.dishID FROM dishes LIMIT $1 OFFSET $2
			)
	`)
	s.getIngredients, err = s.db.Prepare(`
		SELECT ingredients.dishID, products.* FROM dishes
			INNER JOIN ingredients ON ingredients.dishID = dishes.dishID
			INNER JOIN products ON products.productID = ingredients.productID
			WHERE dishes.dishID IN (
				SELECT dishes.dishID FROM dishes LIMIT $1 OFFSET $2
			)
	`)

	if err != nil {
		log.Fatal(err)
	}

	return s
}

func (s *DishService) CreateDishImages(dishImages []models.DishImage) error {
	if len(dishImages) == 0 {
		return nil
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	for _, p := range dishImages {
		_, err := tx.Exec(s.sqlCreateDishImage, p.DishID, p.ImageID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
func (s *DishService) CreateIngredients(ingredients []models.Ingredient) error {
	if len(ingredients) == 0 {
		return nil
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	for _, p := range ingredients {
		_, err := tx.Exec(s.sqlCreateIngredient, p.DishID, p.ProductID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
func (s *DishService) CreateDish(params models.DishParams) (*models.Dish, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.createDish.Exec(
		params.Name,
		params.Description,
		params.Price,
	)
	if err != nil {
		return nil, err
	}

	dishID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var dishImages []models.DishImage
	var ingredients []models.Ingredient

	for _, p := range params.Images {
		dishImages = append(dishImages, models.DishImage{
			DishID:  dishID,
			ImageID: p.ImageID,
		})
	}
	for _, p := range params.Ingredients {
		ingredients = append(ingredients, models.Ingredient{
			DishID:    dishID,
			ProductID: p.ProductID,
		})
	}

	err = s.CreateDishImages(dishImages)
	err = s.CreateIngredients(ingredients)
	if err != nil {
		return nil, err
	}

	return s.GetDishByID(dishID)
}
func (s *DishService) GetDishByID(ID int64) (*models.Dish, error) {
	var dish models.Dish

	err := s.getDishByID.QueryRow(ID).Scan(
		&dish.DishID,
		&dish.Name,
		&dish.Description,
		&dish.Price,
	)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	imageRows, err := s.getDishImageByID.Query(ID)
	if err != nil {
		return nil, err
	}

	for imageRows.Next() {
		image := models.Image{}

		if err := imageRows.Scan(
			&image.ImageID,
			&image.Name,
			&image.Alt,
			&image.URL.Blur,
			&image.URL.Small,
			&image.URL.Medium,
			&image.URL.Large,
		); err != nil {
			return nil, err
		}

		dish.Images = append(dish.Images, image)
	}

	productRows, err := s.getIngredientByID.Query(ID)
	if err != nil {
		return nil, err
	}

	for productRows.Next() {
		product := models.Product{}

		err := productRows.Scan(
			&product.ProductID,
			&product.Name,
		)
		if err != nil {
			return nil, err
		}

		dish.Ingredients = append(dish.Ingredients, product)
	}

	return &dish, nil
}
func (s *DishService) GetDishes(limit int, offset int) ([]models.Dish, error) {
	dishMap := make(map[int64]*models.Dish)

	dishRows, err := s.getDishes.Query(limit, offset)
	if err != nil {
		return nil, err
	}

	for dishRows.Next() {
		dish := models.Dish{}

		if err := dishRows.Scan(
			&dish.DishID,
			&dish.Name,
			&dish.Description,
			&dish.Price,
		); err != nil {
			return nil, err
		}

		dishMap[dish.DishID] = &dish
	}

	imageRows, err := s.getDishImages.Query(limit, offset)
	if err != nil {
		return nil, err
	}

	for imageRows.Next() {
		var dishID int64
		image := models.Image{}

		if err := imageRows.Scan(
			&dishID,
			&image.ImageID,
			&image.Name,
			&image.Alt,
			&image.URL.Blur,
			&image.URL.Small,
			&image.URL.Medium,
			&image.URL.Large,
		); err != nil {
			return nil, err
		}

		dish := dishMap[dishID]
		dish.Images = append(dish.Images, image)
	}

	productRows, err := s.getIngredients.Query(limit, offset)
	if err != nil {
		return nil, err
	}

	for productRows.Next() {
		var dishID int64
		product := models.Product{}

		if err := productRows.Scan(
			&dishID,
			&product.ProductID,
			&product.Name,
		); err != nil {
			return nil, err
		}

		dish := dishMap[dishID]
		dish.Ingredients = append(dish.Ingredients, product)
	}

	var dishes []models.Dish
	for _, dish := range dishMap {
		dishes = append(dishes, *dish)
	}

	return dishes, nil
}

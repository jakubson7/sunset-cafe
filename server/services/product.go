package services

import (
	"database/sql"
	"log"

	"github.com/jakubson7/sunset-cafe/models"
)

type ProductService struct {
	db             *sql.DB
	createProduct  *sql.Stmt
	getAllProducts *sql.Stmt
	updateProduct  *sql.Stmt
	deleteProduct  *sql.Stmt
}

func NewProductService(sqliteService *SqliteService) *ProductService {
	s := &ProductService{}
	var err error

	s.db = sqliteService.DB
	s.createProduct, err = s.db.Prepare(`INSERT INTO products (name) values ($1)`)
	s.getAllProducts, err = s.db.Prepare(`SELECT * FROM products`)
	s.updateProduct, err = s.db.Prepare(`UPDATE products SET name = $2 WHERE productID = $1`)
	s.deleteProduct, err = s.db.Prepare(`
		BEGIN TRANSACTION;
		DELETE FROM ingredients WHERE ingredients.productID = $1;
		DELETE FROM products WHERE products.productID = $1;
		COMMIT;
	`)

	if err != nil {
		log.Fatal(err)
	}

	return s
}

func (s *ProductService) CreateProduct(params models.ProductParams) (*models.Product, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.createProduct.Exec(params.Name)
	if err != nil {
		return nil, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &models.Product{
		ProductID:     ID,
		ProductParams: params,
	}, nil
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	rows, err := s.getAllProducts.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := models.Product{}

		if err := rows.Scan(
			&product.ProductID,
			&product.Name,
		); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (s *ProductService) UpdateProduct(product models.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}

	_, err := s.updateProduct.Exec(product.ProductID, product.Name)
	return err
}

func (s *ProductService) DeleteProduct(ID int64) error {
	_, err := s.deleteProduct.Exec(ID)
	return err
}

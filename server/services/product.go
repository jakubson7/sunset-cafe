package services

import (
	"database/sql"
	"log"

	"github.com/jakubson7/sunset-cafe/models"
)

type ProductService struct {
	db             *sql.DB
	createProduct  *sql.Stmt
	getProductByID *sql.Stmt
	getProducts    *sql.Stmt
	updateProduct  *sql.Stmt
	deleteProduct  *sql.Stmt
}

func NewProductService(sqliteService *SqliteService) *ProductService {
	s := &ProductService{}
	var err error

	s.db = sqliteService.DB
	s.createProduct, err = s.db.Prepare(`INSERT INTO products (name) values ($1)`)
	s.updateProduct, err = s.db.Prepare(`UPDATE products SET name = $2 WHERE productID = $1`)
	s.deleteProduct, err = s.db.Prepare(`DELETE FROM products WHERE productID = $1`)

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

func (s *ProductService) UpdateProduct(product models.Product) (*models.Product, error) {
	if err := product.Validate(); err != nil {
		return nil, err
	}

	_, err := s.updateProduct.Exec(product.ProductID, product.Name)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *ProductService) DeleteProduct(ID int64) error {
	_, err := s.deleteProduct.Exec(ID)
	return err
}

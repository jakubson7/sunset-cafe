package models

import (
	"errors"
	"fmt"
)

type ProductCreate struct {
	Name string `json:"name"`
}

type Product struct {
	ProductID int64 `json:"productID"`
	Timestamp
	ProductCreate
}

const ProductSQL = `
	CREATE TABLE products (
		productID INTEGER,
		createdAt INTEGER NOT NULL,
		updatedAt INTEGER NOT NULL,
		name TEXT NOT NULL,

		PRIMARY KEY (productID)
	)
`

func (create *ProductCreate) Validate() error {
	if isEmpty(create.Name) {
		return create.efrom("Name cannot be an empty string")
	}

	return nil
}

func (create *ProductCreate) efrom(text string) error {
	return errors.New(fmt.Sprintf("(ProductCreate) -> %s", text))
}

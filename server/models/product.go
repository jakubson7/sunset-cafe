package models

import (
	"errors"
)

type ProductParams struct {
	Name string `json:"name"`
}

type Product struct {
	ProductID int64 `json:"productID"`
	ProductParams
}

const ProductSQL = `
	CREATE TABLE products (
		productID INTEGER,
		name TEXT NOT NULL UNIQUE,

		PRIMARY KEY (productID)
	)
`

func (params *ProductParams) Validate() error {
	if isEmpty(params.Name) {
		return errors.New("Name cannot be an empty string")
	}

	return nil
}

package models

import (
	"errors"
	"fmt"
)

type ProductCreate struct {
	Singular string `json:"singular"`
	Plural   string `json:"plural"`
}

type Product struct {
	ProductID int64 `json:"productID"`
	ProductCreate
}

const ProductSQL = `
	CREATE TABLE products (
		productID INTEGER,
		singular TEXT NOT NULL,
		plural TEXT NOT NULL,

		PRIMARY KEY (productID)
	)
`

func (create *ProductCreate) Validate() error {
	if isEmpty(create.Singular) {
		return create.efrom("Name cannot be an empty string")
	}
	if isEmpty(create.Plural) {
		return create.efrom("Name cannot be an empty string")
	}

	return nil
}

func (create *ProductCreate) efrom(text string) error {
	return errors.New(fmt.Sprintf("(ProductCreate) -> %s", text))
}

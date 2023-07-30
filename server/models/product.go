package models

type ProductCreate struct {
	Name string `json:"name"`
}

type Product struct {
	ProductID int64 `json:"productID"`
	Timestamp
	ProductCreate
}

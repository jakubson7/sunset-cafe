package models

import "encoding/json"

type DishCreate struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Dish struct {
	DishID int64 `json:"dishID"`
	Timestamp
	DishCreate
}

const DishSQL = `
	CREATE TABLE dishes (
		dishID INTEGER,
		createdAt INTEGER,
		updatedAt INTEGER,
		name TEXT,
		price REAL,

		PRIMARY KEY (dishID)
	)
`

func (m *Dish) JSON() (string, error) {
	e := newModelError("Dish")
	data, err := json.Marshal(m)
	return string(data), e.Wrap(err)
}

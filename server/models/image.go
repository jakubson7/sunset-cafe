package models

import (
	"database/sql"

	"github.com/jakubson7/sunset-cafe/db"
)

type Image struct {
	ID        int    `json:"ID"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Provider  string `json:"provider"`
	SmallURL  string `json:"smallURL"`
	MediumURL string `json:"mediumURL"`
	BigURL    string `json:"bigURL"`
}

type ImageModel struct {
	db *sql.DB
}

func NewImageModel(db *sql.DB) *ImageModel {
	return &ImageModel{db}
}

func (m *ImageModel) SetupTable() error {
	return db.PrepareAndExec(m.db, `
		CREATE TABLE images (
			imageID INTEGER,
			name TEXT NOT NULL,
			slug TEXT NOT NULL,
			provider TEXT NOT NULL,
			smallURL TEXT NOT NULL,
			mediumURL TEXT NOT NULL,
			bigURL TEXT NOT NULL,

			PRIMARY KEY(imageID)
		)
	`)
}

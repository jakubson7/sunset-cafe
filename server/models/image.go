package models

import (
	"database/sql"
	"log"

	"github.com/gosimple/slug"
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

func NewImage(name string, provider string, smallURL string, mediumURL string, bigURL string) *Image {
	return &Image{
		Name:      name,
		Slug:      slug.Make(name),
		Provider:  provider,
		SmallURL:  smallURL,
		MediumURL: mediumURL,
		BigURL:    bigURL,
	}
}

func NewImageModel(db *sql.DB) *ImageModel {
	return &ImageModel{
		db: db,
	}
}

func (m *ImageModel) SetupTable() {
	err := db.PrepareAndExec(m.db, `
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
	if err != nil {
		log.Fatal(err)
	}
}

func (m *ImageModel) CreateOne(image *Image) error {
	return db.PrepareAndExec(m.db, `
		INSERT INTO images
			(name, slug, provider, smallURL, mediumURL, bigURL)
		VALUES
			($1, $2, $3, $4, $5, $6)
	`, image.Name, image.Slug, image.Provider, image.SmallURL, image.MediumURL, image.BigURL)
}

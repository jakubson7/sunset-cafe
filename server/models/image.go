package models

type ImageParams struct {
	Name string `json:"name"`
	Alt  string `json:"alt"`
}

type ImageVariants struct {
	Blur   string `json:"blur"`
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

type Image struct {
	ImageID int64 `json:"imageID"`
	ImageParams
	Variants ImageVariants `json:"variants"`
}

const ImageSQL = `
	CREATE TABLE images (
		imageID INTEGER,
		name TEXT NOT NULL,
		alt TEXT NOT NULL,
		variants_blur TEXT NOT NULL,
		variants_small TEXT NOT NULL,
		variants_medium TEXT NOT NULL,
		variants_large TEXT NOT NULL,

		PRIMARY KEY (imageID)
	)
`

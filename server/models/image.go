package models

import "errors"

type ImageParams struct {
	Name string `json:"name"`
	Alt  string `json:"alt"`
}

type ImageURL struct {
	Blur   string `json:"blur"`
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

type Image struct {
	ImageID int64 `json:"imageID"`
	ImageParams
	URL ImageURL `json:"URL"`
}

const ImageSQL = `
	CREATE TABLE images (
		imageID INTEGER,
		name TEXT NOT NULL,
		alt TEXT NOT NULL,
		URL_blur TEXT NOT NULL,
		URL_small TEXT NOT NULL,
		URL_medium TEXT NOT NULL,
		URL_large TEXT NOT NULL,

		PRIMARY KEY (imageID)
	)
`

func (params *ImageParams) Validate() error {
	if isEmpty(params.Name) {
		return errors.New("Name cannot be an empty string")
	}
	if isEmpty(params.Alt) {
		return errors.New("Alt cannot be an empty string")
	}

	return nil
}

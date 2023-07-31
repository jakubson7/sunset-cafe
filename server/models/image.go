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
	URL ImageURL `json:"variants"`
}

const ImageSQL = `
	CREATE TABLE images (
		imageID INTEGER,
		name TEXT NOT NULL,
		alt TEXT NOT NULL,
		url_blur TEXT NOT NULL,
		url_small TEXT NOT NULL,
		url_medium TEXT NOT NULL,
		url_large TEXT NOT NULL,

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

package services

import (
	"database/sql"
	"log"

	"github.com/jakubson7/sunset-cafe/models"
)

type ImageService struct {
	db             *sql.DB
	StorageService StorageService
	createImage    *sql.Stmt
	getImageByID   *sql.Stmt
	getImages      *sql.Stmt
	updateImage    *sql.Stmt
	deleteImage    *sql.Stmt
}

func NewImageService(sqliteService *SqliteService, storageService StorageService) *ImageService {
	s := &ImageService{}
	var err error

	s.db = sqliteService.DB
	s.StorageService = storageService
	s.createImage, err = s.db.Prepare(`
		INSERT INTO images (name, alt, URL_blur, URL_small, URL_medium, URL_large)
		VALUES ($1, $2, $3, $4, $5, $6)
	`)
	s.getImageByID, err = s.db.Prepare(`SELECT * FROM images WHERE imageID = $1`)
	s.getImages, err = s.db.Prepare(`SELECT * FROM images LIMIT $1 OFFSET $2`)
	s.updateImage, err = s.db.Prepare(`
		UPDATE images SET
			name = $2, alt = $3, URL_blur = $4, URL_small = $5, URL_medium = $6, URL_large = $7
			WHERE imageID = $1
	`)
	s.deleteImage, err = s.db.Prepare(`
		BEGIN TRANSACTION;
		DELETE FROM dishImages WHERE dishImages.imageID = $1;
		DELETE FROM images WHERE images.imageID = $1;
		COMMIT;
	`)

	if err != nil {
		log.Fatal(err)
	}

	return s
}

func (s *ImageService) CreateImage(data []byte, params models.ImageParams) (*models.Image, error) {
	URL := s.StorageService.GetDefaultImageURL()
	result, err := s.createImage.Exec(
		params.Alt,
		params.Name,
		URL.Blur,
		URL.Small,
		URL.Medium,
		URL.Large,
	)
	if err != nil {
		return nil, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	URL, err = s.StorageService.SaveImage(data, ID)
	if err != nil {
		return nil, err
	}

	image := models.Image{
		ImageID:     ID,
		ImageParams: params,
		URL:         URL,
	}

	err = s.UpdateImage(image)
	if err != nil {
		return nil, err
	}

	return &image, nil
}
func (s *ImageService) GetImageByID(ID int64) (*models.Image, error) {
	image := new(models.Image)

	if err := s.getImageByID.QueryRow().Scan(
		&image.ImageID,
		&image.Name,
		&image.Alt,
		&image.URL.Blur,
		&image.URL.Small,
		&image.URL.Medium,
		&image.URL.Large,
	); err != nil {
		return nil, err
	}

	return image, nil
}
func (s *ImageService) GetImages(limit int, offset int) ([]models.Image, error) {
	rows, err := s.getImages.Query(limit, offset)
	if err != nil {
		return nil, err
	}

	var images []models.Image
	for rows.Next() {
		image := models.Image{}

		if err := rows.Scan(
			&image.ImageID,
			&image.Name,
			&image.Alt,
			&image.URL.Blur,
			&image.URL.Small,
			&image.URL.Medium,
			&image.URL.Large,
		); err != nil {
			return nil, err
		}

		images = append(images, image)
	}

	return images, nil
}

func (s *ImageService) UpdateImage(image models.Image) error {
	_, err := s.updateImage.Exec(
		image.ImageID,
		image.Name,
		image.Alt,
		image.URL.Blur,
		image.URL.Small,
		image.URL.Medium,
		image.URL.Large,
	)
	return err
}

func (s *ImageService) DeleteImage(ID int64) error {
	_, err := s.deleteImage.Exec(ID)
	return err
}

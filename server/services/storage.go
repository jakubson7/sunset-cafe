package services

import "github.com/jakubson7/sunset-cafe/models"

type StorageService[Config any] interface {
	Init(config *Config)
	GetDefaultImageURL() models.ImageURL
	SaveImage(data []byte, ID int64) (*models.ImageURL, error)
	GetImageURL(ID int64) (*models.ImageURL, error)
	DeleteImage(ID int64) error
}

type LocalStorageServiceConfig struct {
	addr string
}

type LocalStorageService struct {
	defaultImageURL models.ImageURL
}

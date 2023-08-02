package services

import (
	"log"

	"github.com/jakubson7/sunset-cafe/models"
	"github.com/jakubson7/sunset-cafe/utils"
)

type StorageService interface {
	Init()
	GetDefaultImageURL() models.ImageURL
	SaveImage(data []byte, ID int64) (models.ImageURL, error)
	GetImageURL(ID int64) (models.ImageURL, error)
	DeleteImage(ID int64) error
}

type LocalStorageServiceConfig struct {
	addr    string
	dirname string
	quality int
}

type LocalStorageService struct {
	defaultImageURL models.ImageURL
	config          LocalStorageServiceConfig
}

func NewLocalStorageService(config LocalStorageServiceConfig) *LocalStorageService {
	s := new(LocalStorageService)

	err := utils.CreateFolder(config.dirname)
	if err != nil {
		log.Fatal(err)
	}

	s.config = config

	return s
}

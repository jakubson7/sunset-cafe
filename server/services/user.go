package services

import (
	"github.com/jakubson7/sunset-cafe/models"
)

type UserService[T any] interface {
	CreateUser(create models.UserCreate) (models.User, error)
	GetUserByID(ID int) (models.User, error)
	GetUsers(limit int, offset int) ([]models.User, error)
}

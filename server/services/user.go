package services

import (
	"github.com/jakubson7/sunset-cafe/models"
)

type UserService[T any] interface {
	GetUserByID(ID int) (models.User, error)
	CreateUser(create models.UserCreate) (models.User, error)
}

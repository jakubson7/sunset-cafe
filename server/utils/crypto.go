package utils

import (
	"github.com/jakubson7/sunset-cafe/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	salt := config.USER_PASSWORD_SALT
	data := []byte(salt + password)

	hashedData, err := bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost)
	return string(hashedData), err
}
func VerfiyPassword(hash, password string) error {
	salt := config.USER_PASSWORD_SALT
	data := []byte(salt + password)
	return bcrypt.CompareHashAndPassword([]byte(password), data)
}

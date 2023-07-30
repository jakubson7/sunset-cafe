package models

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jakubson7/sunset-cafe/config"
	"golang.org/x/crypto/bcrypt"
)

type UserCreate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type User struct {
	UserID int64 `json:"userID"`
	Timestamp
	UserCreate
}

const UserSQL = `
	CREATE TABLE users (
		userID INTEGER,
		createdAt INTEGER NOT NULL,
		updatedAt INTEGER NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		name TEXT NOT NULL,

		PRIMARY KEY (userID)
	)
`

func (create *UserCreate) Validate() error {
	if trim(create.Email) == "" {
		return create.efrom("Email cannot be an empty string")
	}

	if isEmpty(create.Password) {
		return create.efrom("Password cannot be an empty string")
	} else if len(trim(create.Password)) < 8 {
		return create.efrom("Password has to be at least 8 characters long")
	}

	if isEmpty(create.Name) {
		return create.efrom("Name cannot be an empty string")
	}

	return nil
}

func (m *User) Validate() error {
	if err := m.Timestamp.Validate(); err != nil {
		return m.ewrap(err)
	}
	if err := m.UserCreate.Validate(); err != nil {
		return m.ewrap(err)
	}

	return nil
}

func (create *UserCreate) HashPassword() error {
	salt := config.USER_PASSWORD_SALT
	data := []byte(salt + create.Password)

	hashedData, err := bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost)
	if err != nil {
		return create.ewrap(err)
	}

	create.Password = string(hashedData)
	return nil
}
func (create *UserCreate) VerfiyPassword(password string) error {
	salt := config.USER_PASSWORD_SALT
	data := []byte(salt + password)
	err := bcrypt.CompareHashAndPassword([]byte(create.Password), data)
	return create.ewrap(err)
}

func (m *User) JSON() (string, error) {
	data, err := json.Marshal(m)
	return string(data), m.ewrap(err)
}

func (create *UserCreate) efrom(text string) error {
	return errors.New(fmt.Sprintf("(UserCreate) -> %s", text))
}
func (create *UserCreate) ewrap(err error) error {
	if err == nil {
		return nil
	}
	return errors.New(fmt.Sprintf("(UserCreate) -> %v", err))
}

func (m *User) ewrap(err error) error {
	if err == nil {
		return nil
	}
	return errors.New(fmt.Sprintf("(User) -> %v", err))
}

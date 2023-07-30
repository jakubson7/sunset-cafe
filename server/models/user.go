package models

import (
	"encoding/json"
	"errors"
	"fmt"
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
		createdAt INTEGER,
		updatedAt INTEGER,
		email TEXT,
		password TEXT,
		name TEXT,

		PRIMARY KEY (userID)
	)
`

func (create *UserCreate) efrom(text string) error {
	return errors.New(fmt.Sprintf("(UserCreate) -> %s", text))
}
func (create *UserCreate) ewrap(err error) error {
	return errors.New(fmt.Sprintf("(UserCreate) -> %v", err))
}

func (create *UserCreate) Validate() error {
	if trim(create.Email) == "" {
		return create.efrom("Email cannot be an empty string")
	}

	if trim(create.Password) == "" {
		return create.efrom("Password cannot be an empty string")
	} else if len(trim(create.Password)) < 8 {
		return create.efrom("Password has to be at least 8 characters long")
	}

	if trim(create.Name) == "" {
		return create.efrom("Name cannot be an empty string")
	}

	return nil
}

func (create *UserCreate) HashPassword() error {
	return nil
}

func (m *User) Validate() error {
	if err := m.Timestamp.Validate(); err != nil {
		return err
	}
	if err := m.UserCreate.Validate(); err != nil {
		return err
	}

	return nil
}

func (m *User) JSON() (string, error) {
	e := newModelError("User")
	data, err := json.Marshal(m)
	return string(data), e.Wrap(err)
}

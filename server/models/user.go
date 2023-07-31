package models

import (
	"errors"
)

type UserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type User struct {
	UserID int64 `json:"userID"`
	UserParams
}

const UserSQL = `
	CREATE TABLE users (
		userID INTEGER,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		name TEXT NOT NULL,

		PRIMARY KEY (userID)
	)
`

func (params *UserParams) Validate() error {
	if trim(params.Email) == "" {
		return errors.New("Email cannot be an empty string")
	}

	if isEmpty(params.Password) {
		return errors.New("Password cannot be an empty string")
	} else if len(trim(params.Password)) < 8 {
		return errors.New("Password has to be at least 8 characters long")
	}

	if isEmpty(params.Name) {
		return errors.New("Name cannot be an empty string")
	}

	return nil
}

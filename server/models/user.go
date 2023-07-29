package models

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

func (create *UserCreate) Validate() error {
	e := newValidationError("UserCreate")

	if trim(create.Email) == "" {
		return e.From("Email cannot be an empty string")
	}

	if trim(create.Password) == "" {
		return e.From("Password cannot be an empty string")
	} else if len(trim(create.Password)) < 8 {
		return e.From("Password has to be at least 8 characters long")
	}

	if trim(create.Name) == "" {
		return e.From("Name cannot be an empty string")
	}

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

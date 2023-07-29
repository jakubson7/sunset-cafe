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
	e := newValidationError("User")

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

func NewUser(create UserCreate, timestamp Timestamp, ID int64) (User, error) {
	user := User{}

	if err := create.Validate(); err != nil {
		return user, err
	}

	user.UserCreate = create
	user.UserID = ID
	user.Timestamp = timestamp

	return user, nil
}

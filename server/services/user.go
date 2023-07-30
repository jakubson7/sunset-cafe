package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jakubson7/sunset-cafe/models"
)

type UserService struct {
	name        string
	db          *sql.DB
	createUser  *sql.Stmt
	getUserByID *sql.Stmt
	getUsers    *sql.Stmt
}

func NewSqliteUserService(sqliteService *SqliteService) *UserService {
	s := &UserService{}
	var err error

	s.db = sqliteService.DB
	s.createUser, err = s.db.Prepare(`INSERT INTO users (createdAt, updatedAt, email, password, name) VALUES ($1, $2, $3, $4, $5)`)
	s.getUserByID, err = s.db.Prepare(`SELECT * FROM users WHERE userID = $1`)
	s.getUsers, err = s.db.Prepare(`SELECT * FROM users LIMIT $1 OFFSET $2`)

	if err != nil {
		s.efatal(err)
	}

	return s
}

func (s *UserService) CreateUser(create models.UserCreate) (models.User, error) {
	ts := models.NewTimestamp()

	if err := create.Validate(); err != nil {
		return models.User{}, s.ewrap(err)
	}

	result, err := s.createUser.Exec(
		ts.CreatedAt,
		ts.UpdatedAt,
		create.Email,
		create.Password,
		create.Name,
	)
	if err != nil {
		return models.User{}, s.ewrap(err)
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return models.User{}, s.ewrap(err)
	}

	return models.User{
		UserID:     ID,
		Timestamp:  ts,
		UserCreate: create,
	}, s.ewrap(err)
}

func (s *UserService) GetUserByID(ID int64) (models.User, error) {
	user := models.User{}

	err := s.getUserByID.QueryRow(ID).Scan(
		&user.UserID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
		&user.Password,
		&user.Name,
	)
	if err != nil {
		return user, s.ewrap(err)
	}

	err = user.Validate()
	return user, s.ewrap(err)
}

func (s *UserService) GetUsers(limit int, offset int) ([]models.User, error) {
	rows, err := s.getUsers.Query(limit, offset)
	if err != nil {
		return nil, s.ewrap(err)
	}

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(
			&user.UserID,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Email,
			&user.Password,
			&user.Name,
		)
		if err != nil {
			return nil, s.ewrap(err)
		}
		if err := user.Validate(); err != nil {
			return nil, s.ewrap(err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *UserService) ewrap(err error) error {
	if err == nil {
		return nil
	}
	return errors.New(fmt.Sprintf("(UserService) -> %v", err))
}
func (s *UserService) efatal(err error) {
	log.Fatal(s.ewrap(err))
}

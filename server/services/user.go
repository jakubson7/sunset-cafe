package services

import (
	"database/sql"
	"log"

	"github.com/jakubson7/sunset-cafe/models"
	"github.com/jakubson7/sunset-cafe/utils"
)

type UserService struct {
	db          *sql.DB
	createUser  *sql.Stmt
	getUserByID *sql.Stmt
	getUsers    *sql.Stmt
}

func NewUserService(sqliteService *SqliteService) *UserService {
	s := &UserService{}
	var err error

	s.db = sqliteService.DB
	s.createUser, err = s.db.Prepare(`INSERT INTO users (createdAt, updatedAt, email, password, name) VALUES ($1, $2, $3, $4, $5)`)
	s.getUserByID, err = s.db.Prepare(`SELECT * FROM users WHERE userID = $1`)
	s.getUsers, err = s.db.Prepare(`SELECT * FROM users LIMIT $1 OFFSET $2`)

	if err != nil {
		log.Fatal(err)
	}

	return s
}

func (s *UserService) CreateUser(params models.UserParams) (*models.User, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(params.Password)
	if err != nil {
		return nil, err
	}

	params.Password = hashedPassword

	result, err := s.createUser.Exec(
		params.Email,
		params.Password,
		params.Name,
	)
	if err != nil {
		return nil, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &models.User{
		UserID:     ID,
		UserParams: params,
	}, err
}

func (s *UserService) GetUserByID(ID int64) (*models.User, error) {
	user := &models.User{}

	err := s.getUserByID.QueryRow(ID).Scan(
		&user.UserID,
		&user.Email,
		&user.Password,
		&user.Name,
	)
	if err != nil {
		return nil, err
	}

	err = user.Validate()
	return user, err
}

func (s *UserService) GetUsers(limit int, offset int) ([]models.User, error) {
	rows, err := s.getUsers.Query(limit, offset)
	if err != nil {
		return nil, err
	}

	users := []models.User{}
	for rows.Next() {
		user := models.User{}

		if err := rows.Scan(
			&user.UserID,
			&user.Email,
			&user.Password,
			&user.Name,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

package services

import (
	"database/sql"

	"github.com/jakubson7/sunset-cafe/models"
)

type SqliteUserService struct {
	name                string
	db                  *sql.DB
	sql_createUserTable string
	sql_createUser      string
	sql_getUserByID     string
	sql_getUsers        string
	stmt_createUser     *sql.Stmt
	stmt_getUserByID    *sql.Stmt
	stmt_getUsers       *sql.Stmt
}

func NewSqliteUserService(db *sql.DB) (*SqliteUserService, error) {
	e := newServiceError("SqliteUserService", "NewSqliteUserService")
	s := &SqliteUserService{}

	s.name = "SqliteUserService"
	s.db = db
	s.sql_createUserTable = `
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
	s.sql_createUser = `
		INSERT INTO users
			(createdAt, updatedAt, email, password, name)
		VALUES
			($1, $2, $3, $4, $5);	
	`
	s.sql_getUserByID = `SELECT * FROM users WHERE userID = $1`
	s.sql_getUsers = `SELECT * FROM users LIMIT $1 OFFSET $2`

	_, err := db.Exec(s.sql_createUserTable)
	if err != nil {
		return nil, e.Wrap(err)
	}

	s.stmt_createUser, err = s.db.Prepare(s.sql_createUser)
	s.stmt_getUserByID, err = s.db.Prepare(s.sql_getUserByID)
	s.stmt_getUsers, err = s.db.Prepare(s.sql_getUsers)

	return s, e.Wrap(err)
}

func (s *SqliteUserService) CreateUser(create models.UserCreate) (models.User, error) {
	e := newServiceError(s.name, "CreateUser")
	ts := models.NewTimestamp()

	if err := create.Validate(); err != nil {
		return models.User{}, e.Wrap(err)
	}

	result, err := s.stmt_createUser.Exec(
		ts.CreatedAt,
		ts.UpdatedAt,
		create.Email,
		create.Password,
		create.Name,
	)
	if err != nil {
		return models.User{}, e.Wrap(err)
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return models.User{}, e.Wrap(err)
	}

	return models.User{
		UserID:     ID,
		Timestamp:  ts,
		UserCreate: create,
	}, e.Wrap(err)
}

func (s *SqliteUserService) GetUserByID(ID int64) (models.User, error) {
	e := newServiceError(s.name, "GetUserByID")
	user := models.User{}

	err := s.stmt_getUserByID.QueryRow(ID).Scan(
		&user.UserID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
		&user.Password,
		&user.Name,
	)
	if err != nil {
		return user, e.Wrap(err)
	}

	err = user.Validate()
	return user, e.Wrap(err)
}

func (s *SqliteUserService) GetUsers(limit int, offset int) ([]models.User, error) {
	e := newServiceError(s.name, "GetUsers")

	rows, err := s.stmt_getUsers.Query(limit, offset)
	if err != nil {
		return nil, e.Wrap(err)
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
			return nil, e.Wrap(err)
		}
		if err := user.Validate(); err != nil {
			return nil, e.Wrap(err)
		}

		users = append(users, user)
	}

	return users, nil
}

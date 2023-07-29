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
	stmt_createUser     *sql.Stmt
	stmt_getUserByID    *sql.Stmt
}

func NewSqliteUserService(db *sql.DB) (*SqliteUserService, error) {
	e := newServiceError("SqliteUserService", "NewSqliteUserService")
	s := &SqliteUserService{}

	s.name = "SqliteUserService"
	s.db = db
	s.sql_createUserTable = `
		CREATE TABLE users (
			userID INTEGER,
			createdAt TIMESTAMP,
			updatedAt TIMESTAMP,
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

	_, err := db.Exec(s.sql_createUserTable)
	if err != nil {
		return nil, e.Wrap(err)
	}

	s.stmt_createUser, err = s.db.Prepare(s.sql_createUser)
	s.stmt_getUserByID, err = s.db.Prepare(s.sql_getUserByID)

	return s, e.Wrap(err)
}

func (s *SqliteUserService) CreateUser(create models.UserCreate) (models.User, error) {
	e := newServiceError(s.name, "CreateUser")
	ts := models.NewTimestamp()

	result, err := s.stmt_createUser.Exec(ts.CreatedAt, ts.UpdatedAt, create.Email, create.Password, create.Name)
	if err != nil {
		return models.User{}, e.Wrap(err)
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return models.User{}, e.Wrap(err)
	}

	user, err := models.NewUser(create, ts, ID)
	return user, e.Wrap(err)
}

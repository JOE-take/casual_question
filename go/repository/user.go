package repository

import (
	"casual_question/models"
	"database/sql"
	"errors"
)

type UserRepositorier interface {
	Create(u *models.User) error
	ReadByEmail(i *models.User) (*models.User, error)
}

type UserRepository struct {
	repo *sql.DB
}

func NewUserRepository(repo *sql.DB) *UserRepository {
	return &UserRepository{repo: repo}
}

func (r UserRepository) Create(u *models.User) error {
	db := r.repo

	insert, err := db.Prepare("insert into Users values(?, ?, ?, ?)")
	if err != nil {
		return err
	}

	result, err := insert.Exec(u.UserID, u.UserName, u.Email, u.Password)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return errors.New("user already exist")
	}
	return nil
}

func (r UserRepository) ReadByEmail(u *models.User) (*models.User, error) {
	db := r.repo
	result := &models.User{}

	row := db.QueryRow("select * from Users where email = ?", u.Email)
	if err := row.Scan(result); err != nil {
		return nil, err
	}

	return result, nil
}

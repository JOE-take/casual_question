package repository

import (
	"casual_question/models"
	"database/sql"
	"errors"
)

type RefTokenRepositorier interface {
	Create(t *models.RefreshToken) error
	Delete(token string) error
	ReadByToken(token string) (*models.RefreshToken, error)
}

type RefRepository struct {
	repo *sql.DB
}

func NewRefRepository(repo *sql.DB) *RefRepository {
	return &RefRepository{repo: repo}
}

func (r RefRepository) Create(t *models.RefreshToken) error {
	db := r.repo

	insert, err := db.Prepare("insert into RefreshTokens (token, user_id, expiry) values (?, ?, FROM_UNIXTIME(?))")
	if err != nil {
		return err
	}

	_, err = insert.Exec(t.Token, t.UserID, t.Expiry)
	if err != nil {
		return err
	}

	return nil
}

func (r RefRepository) Delete(token string) error {
	db := r.repo
	delete, err := db.Prepare("delete from RefreshTokens where token = ?")
	if err != nil {
		return err
	}

	result, err := delete.Exec(token)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		err := errors.New("no matching token found")
		return err
	}

	return nil
}

func (r RefRepository) ReadByToken(token string) (*models.RefreshToken, error) {
	db := r.repo
	result := &models.RefreshToken{}

	row := db.QueryRow("select token, user_id, UNIX_TIMESTAMP(expiry) from RefreshTokens where token = ?", token)
	err := row.Scan(&result.Token, &result.UserID, &result.Expiry)
	if err != nil {
		return nil, err
	}

	return result, nil
}

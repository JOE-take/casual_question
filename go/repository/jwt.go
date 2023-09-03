package repository

import (
	"casual_question/models"
	"database/sql"
)

type RefTokenRepositorier interface {
	Create(t *models.RefreshToken) error
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

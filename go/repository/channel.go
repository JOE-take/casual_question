package repository

import (
	"casual_question/models"
	"database/sql"
	"errors"
)

type ChannelRepositorier interface {
	Create(c *models.Channel) error
}

type ChannelRepository struct {
	repo *sql.DB
}

func NewChannelRepository(repo *sql.DB) *ChannelRepository {
	return &ChannelRepository{repo: repo}
}

func (r ChannelRepository) Create(c *models.Channel) error {
	db := r.repo

	insert, err := db.Prepare("insert into Channels values (?, ?)")
	if err != nil {
		return err
	}

	result, err := insert.Exec(c.ChannelID, c.Owner)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return errors.New("no user created")
	}

	return nil
}

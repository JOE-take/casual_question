package repository

import (
	"casual_question/models"
	"database/sql"
)

type QuestionRepositorier interface {
	Create(q *models.Question) error
}

type QuestionRepository struct {
	repo *sql.DB
}

func NewQuestionRepository(repo *sql.DB) *QuestionRepository {
	return &QuestionRepository{repo: repo}
}

func (r QuestionRepository) Create(q *models.Question) error {
	db := r.repo
	insert, err := db.Prepare("insert into Questions values (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = insert.Exec(q.ChannelID, q.ID, q.Content)
	if err != nil {
		return err
	}

	return nil
}

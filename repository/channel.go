package repository

import (
	"casual_question/models"
	"database/sql"
	"errors"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ChannelRepositorier interface {
	CreateUnique(string) (string, error)
	ReadAllByID(string) ([]models.Question, error)
}

type ChannelRepository struct {
	repo *sql.DB
}

func NewChannelRepository(repo *sql.DB) *ChannelRepository {
	return &ChannelRepository{repo: repo}
}

func (r ChannelRepository) CreateUnique(owner string) (string, error) {
	db := r.repo
	channel := &models.Channel{}

	// 新しいIDの生成
	newID, err := createUniqueID(db)
	if err != nil {
		return "", err
	}

	//チャンネル情報の決定
	channel.ChannelID = newID
	channel.Owner = owner

	insert, err := db.Prepare("insert into Channels (channel_id, owner) values (?, ?)")
	if err != nil {
		return "", err
	}
	defer insert.Close()

	result, err := insert.Exec(channel.ChannelID, channel.Owner)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	} else if rowsAffected == 0 {
		return "", errors.New("no channel made")
	}

	return newID, nil
}

func (r ChannelRepository) ReadAllByID(channelID string) ([]models.Question, error) {
	db := r.repo
	var result []models.Question

	rows, err := db.Query("select channel_id, id, content, created_at from Questions where channel_id = ?", channelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 1レコードずつスキャンしてresultに追加
	for rows.Next() {
		tmp := models.Question{}
		var createdAtStr string
		err := rows.Scan(&tmp.ChannelID, &tmp.ID, &tmp.Content, &createdAtStr)
		if err != nil {
			return nil, err
		}

		tmp.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr) // 日時のフォーマットを適切に指定
		result = append(result, tmp)
	}

	return result, nil
}

func createUniqueID(db *sql.DB) (string, error) {

	// 新しいIDが見つかるまで回す タイムアウトを実装するべき？
	for {
		seed := time.Now().UnixNano()
		random := rand.New(rand.NewSource(seed))
		newID := strconv.Itoa(random.Intn(1000))

		var existingID string
		row := db.QueryRow("select * from Channels where channel_id = ?", newID)

		// 該当するレコードが存在しなければこれが新しいID
		err := row.Scan(&existingID)
		if err == sql.ErrNoRows {
			return newID, nil
		}

		if err != nil {
			return "0", err
		}
	}
}
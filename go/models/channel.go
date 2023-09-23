package models

import "time"

type Channel struct {
	ChannelID string `db:"channel_id"`
	Owner     string `db:"owner"`
}

type Question struct {
	ID        string `db:"id"`
	ChannelID string `db:"channel_id"`
	Content   string `db:"content"`
	CreatedAt time.Time
}

package models

type RefreshToken struct {
	Token  string `db:"token"`
	UserID string `db:"user_id"`
	Expiry int64  `db:"expiry"`
}

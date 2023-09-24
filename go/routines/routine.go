package routines

import (
	"database/sql"
	"log"
	"time"
)

type RoutineInfo struct {
	DB *sql.DB
}

func ControlRoutines(info RoutineInfo) {
	for {
		time.Sleep(12 * time.Hour)
		deleteExpiredChannels(info.DB)
		deleteExpiredRefreshTokens(info.DB)
	}
}

// 24時間経過したチャンネルを削除
func deleteExpiredChannels(db *sql.DB) {
	delete, err := db.Prepare("DELETE FROM Channels WHERE TIMESTAMPDIFF(HOUR, createdAt, NOW()) >= 24")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = delete.Exec(time.Now())
	if err != nil {
		log.Println(err)
		return
	}
}

func deleteExpiredRefreshTokens(db *sql.DB) {
	delete, err := db.Prepare("delete from RefreshTokens where expiry < NOW()")
	if err != nil {
		log.Println(err)
		return
	}

	_, err = delete.Exec()
	if err != nil {
		log.Println(err)
		return
	}
}

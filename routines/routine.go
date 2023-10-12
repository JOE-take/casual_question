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
		time.Sleep(2 * time.Hour)
		deleteExpiredChannels(info.DB)
		deleteExpiredRefreshTokens(info.DB)
	}
}

// 24時間経過したチャンネルを削除
func deleteExpiredChannels(db *sql.DB) {
	deletePrepare, err := db.Prepare("DELETE FROM Channels WHERE TIMESTAMPDIFF(HOUR, created_at, NOW()) >= 24")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = deletePrepare.Exec()
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

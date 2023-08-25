package db

import (
	"database/sql"
	"log"
)

func InitializeDB() *sql.DB {

	db, err := sql.Open("mysql", "root:password@tcp(CQ-db:3306)/cq")
	if err != nil {
		log.Fatal("failed to open database")
	}

	if err := db.Ping(); err != nil {
		log.Fatal("failed to connect database")
	}

	return db
}

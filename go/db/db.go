package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitializeDB() *sql.DB {

	db, err := sql.Open("mysql", "root:password@tcp(CQ-db:3306)/cq")
	if err != nil {
		log.Fatal("failed to open database: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("failed to connect database", err)
	}

	return db
}

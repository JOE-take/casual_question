package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

const (
	MaxRetry = 5
	WaitTime = 10
)

func InitializeDB() *sql.DB {

	db, err := openDB()
	if err != nil {
		log.Fatal("failed to open database.\n", err)
	}

	err = pingDB(db)
	if err != nil {
		log.Fatal("failed to connect database.\n", err)
	}

	return db
}

func openDB() (*sql.DB, error) {
	var db *sql.DB
	var err error

	dbUser := os.Getenv("CLOUD_SQL_USER")
	dbPass := os.Getenv("CLOUD_SQL_PASS")
	cloudName := os.Getenv("CLOUD_SQL_NAME")
	dbName := os.Getenv("CLOUD_SQL_DBNAME")
	dbSource := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=true", dbUser, dbPass, cloudName, dbName)

	for i := 0; i <= MaxRetry; i++ {
		db, err = sql.Open("mysql", dbSource)
		if err == nil {
			break
		}
		time.Sleep(WaitTime * time.Second)
	}

	return db, err
}

func pingDB(db *sql.DB) error {
	var err error

	for i := 0; i <= MaxRetry; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(WaitTime * time.Second)
	}

	return err
}

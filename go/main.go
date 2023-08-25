package main

import (
	"casual_question/db"
	"casual_question/router"
)

func main() {
	db := db.InitializeDB()
	r := router.NewRouter(db)
	r.Run(":8080")
}

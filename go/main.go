package main

import (
	"casual_question/db"
	"casual_question/router"
	"casual_question/routines"
	"time"
)

func main() {
	dbCon := db.InitializeDB()
	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)
	go routines.ControlRoutines(routines.RoutineInfo{DB: dbCon})
	r := router.NewRouter(dbCon)
	r.Run(":8080")
}

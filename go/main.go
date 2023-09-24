package main

import (
	"casual_question/db"
	"casual_question/router"
	"casual_question/routines"
)

func main() {
	dbCon := db.InitializeDB()
	go routines.ControlRoutines(routines.RoutineInfo{DB: dbCon})
	r := router.NewRouter(dbCon)
	r.Run(":8080")
}

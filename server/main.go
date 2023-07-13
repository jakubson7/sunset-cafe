package main

import (
	"github.com/jakubson7/sunset-cafe/db"
	"github.com/jakubson7/sunset-cafe/repos"
	"github.com/jakubson7/sunset-cafe/router"
)

func main() {
	sqlite := db.NewSqlite()
	defer sqlite.Close()

	repos.Setup(sqlite)
	repos.Refresh()
	repos.MockMeals()

	r := router.NewRouter()
	router.Serve(r)
}

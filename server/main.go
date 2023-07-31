package main

import "github.com/jakubson7/sunset-cafe/services"

func main() {
	sqlite := services.NewMemorySqliteService()
	sqlite.CreateTables()
}

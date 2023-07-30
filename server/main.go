package main

import (
	"fmt"

	"github.com/jakubson7/sunset-cafe/models"
	"github.com/jakubson7/sunset-cafe/services"
)

func main() {
	sqliteService := services.NewMemorySqliteService()
	sqliteService.CreateTables()

	userService := services.NewUserService(sqliteService)

	uc := models.UserCreate{"test@gmail.com", "12341234", "James"}
	u, err := userService.CreateUser(uc)
	fmt.Printf("%v \n %v", u, err)
	fmt.Printf("\n---\n %v", u.VerfiyPassword("12341134"))
}

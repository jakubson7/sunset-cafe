package main

import (
	"fmt"
	"log"

	"github.com/jakubson7/sunset-cafe/db"
	"github.com/jakubson7/sunset-cafe/models"
	"github.com/jakubson7/sunset-cafe/services"
)

func main() {
	db := db.NewMemorySqlite()
	userService, err := services.NewSqliteUserService(db)
	if err != nil {
		log.Fatal(err)
	}

	user, err := userService.CreateUser(models.UserCreate{"kuba@gmail.com", "asd123123", "adsasd"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", user)
}

package main

import (
	"fmt"
	"log"

	"github.com/jakubson7/sunset-cafe/models"
	"github.com/jakubson7/sunset-cafe/services"
)

func main() {
	db, err := services.NewMemorySqliteService()
	if err != nil {
		log.Fatal(err)
	}

	userService, err := services.NewSqliteUserService(db)
	if err != nil {
		log.Fatal(err)
	}

	_, err = userService.CreateUser(models.UserCreate{"kuba@gmail.com", "2312asd333", "adsasd"})
	_, err = userService.CreateUser(models.UserCreate{"kuba@gmail.com", "2312asd333", "adsasd"})
	if err != nil {
		log.Fatal(err)
	}

	user, err := userService.GetUserByID(1)
	fmt.Printf("%v --- %v \n", err, user)
	users, err := userService.GetUsers(10, 1)
	fmt.Printf("%v --- %v \n", err, users)
}

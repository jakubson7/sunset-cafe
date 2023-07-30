package main

import (
	"fmt"

	"github.com/jakubson7/sunset-cafe/models"
)

func main() {
	c := models.UserCreate{"abc@gmail.com", "12345678", "Patryk"}
	ts := models.NewTimestamp()
	user := models.User{1, ts, c}

	fmt.Println(user)
	fmt.Println(user.Validate())
	fmt.Println(user.HashPassword())
	fmt.Println(user)
	fmt.Println(user.VerfiyPassword("12345d78"))
	fmt.Println(user.VerfiyPassword("12345678"))
}

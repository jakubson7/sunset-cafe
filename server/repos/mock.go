package repos

import (
	"log"

	"github.com/jakubson7/sunset-cafe/lib"
)

func MockMeals() {
	meal := lib.Meal{
		Name:  "Spaghetti",
		Slug:  "spaghetti",
		Price: 20,
		ImgID: 1,
	}

	for i := 1; i < 10; i++ {
		err := Meal.CreateOne(&meal)
		if err != nil {
			log.Fatal(err)
		}
	}
}

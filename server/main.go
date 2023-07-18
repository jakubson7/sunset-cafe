package main

import "github.com/jakubson7/sunset-cafe/api"

func main() {
	app := api.NewAPI()
	app.Start()
}

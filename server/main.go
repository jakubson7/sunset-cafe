package main

import "github.com/jakubson7/sunset-cafe/api"

func main() {
	sunsetCafeApi := api.NewApi()
	sunsetCafeApi.Mock()
	sunsetCafeApi.Start()
}

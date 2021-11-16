package main

import (
	"restApi/config"
	"restApi/routes"
)

func main() {
	config.InitDB()

	e := routes.New()

	e.Start(":8000")
}

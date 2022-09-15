package main

import (
	"crud-dynamic/config"
	"crud-dynamic/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}

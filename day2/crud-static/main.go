package main

import (
	"crud-static/routes"
)

func main() {

	e := routes.New()

	// start the server, and log if it fails
	e.Start(":8080")
}

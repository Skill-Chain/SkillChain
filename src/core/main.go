package main

import (
	"APImod/src/presentation-layer/api/routes"
	"log"
)

func main() {
	err := routes.Router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err.Error())
	}

	routes.GetRoutes()
}

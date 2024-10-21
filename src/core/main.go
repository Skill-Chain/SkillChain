package main

import (
	"APImod/src/infrastructure/database"
	"APImod/src/presentation-layer/api/routes"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := routes.Router.Run(port)
	if err != nil {
		log.Fatal(err.Error())
	}

	routes.GetRoutes()
	database.ConnectDB()
}

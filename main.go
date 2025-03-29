package main

import (
	"log"
	"net/http"
	"os"

	"github.com/johneliud/authentication_project/backend/database"
	"github.com/johneliud/authentication_project/backend/routes"
	"github.com/johneliud/authentication_project/backend/utils"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) > 2 {
		log.Println("Usage: go run main.go OR go run main.go [PORT]")
		return
	}

	_ = os.Mkdir("data", 0o700)

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Failed to load .env file: %v. Using default values\n", err)
	}

	database.InitDB()

	defer func() {
		if database.DB != nil {
			database.DB.Close()
		}
	}()

	routes.InitRoutes()

	port, err := utils.GetPort()
	if err != nil {
		log.Printf("Failed to get port: %v\n", err)
		return
	}

	log.Printf("Starting server on port %s\n", port)
	if err = http.ListenAndServe(port, nil); err != nil {
		log.Printf("Failed to start server: %v\n", err)
		return
	}
}

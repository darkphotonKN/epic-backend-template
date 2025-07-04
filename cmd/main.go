package main

import (
	"fmt"
	"log"
	"os"

	"github.com/darkphotonKN/eh-hub-data-orchestration-platform/config"
	"github.com/joho/godotenv"
)

/**
* Main entry point to entire application.
* NOTE: Keep code here as clean and minimal as possible.
**/
func main() {
	// env setup
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// router setup
	router := config.SetupRouter()

	defaultDevPort := ":8080"

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultDevPort
	}

	// starts server and listen on port
	router.Run(fmt.Sprintf(":%s", port)) // port = ":" + PORT
}

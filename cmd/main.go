package main

import (
	"fmt"
	"log"

	"github.com/darkphotonKN/go-figma-mcp/config"
	_ "github.com/joho/godotenv/autoload" // auto-load env vars
)

func main() {
	// Load configuration
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Setup router
	router := config.SetupRouter(appConfig)

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)

	if err := router.Run(port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

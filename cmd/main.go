package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/darkphotonKN/go-figma-mcp/config"
	_ "github.com/joho/godotenv" // load env vars
)

func main() {
	// load configuration
	err := config.LoadConfig()

	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}


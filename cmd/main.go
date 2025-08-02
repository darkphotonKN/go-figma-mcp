package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/darkphotonKN/go-figma-mcp/config"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// TODO: Initialize MCP server
	// TODO: Register tools and resources
	// TODO: Start server

	log.Printf("Starting %s v%s", cfg.Server.Name, cfg.Server.Version)
	log.Println("MCP server implementation to be added")
}
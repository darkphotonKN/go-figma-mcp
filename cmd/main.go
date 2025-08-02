package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/darkphotonKN/go-figma-mcp/config"
	"github.com/darkphotonKN/go-figma-mcp/internal/figma"
	"github.com/darkphotonKN/go-figma-mcp/pkg/mcp"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Create MCP server capabilities
	capabilities := mcp.NewCapabilitiesBuilder().
		EnableTools().
		EnableResources(false).
		Build()

	// Create MCP server
	server := mcp.NewServer(mcp.ServerConfig{
		Name:         cfg.Server.Name,
		Version:      cfg.Server.Version,
		Capabilities: capabilities,
	})

	// Register Figma tools
	figmaService := figma.NewService(cfg.Figma.APIKey)
	if err := figma.RegisterTools(server, figmaService); err != nil {
		log.Fatal("Failed to register Figma tools:", err)
	}

	// Setup graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Listen for interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Received interrupt signal, shutting down...")
		cancel()
	}()

	// Start the MCP server
	log.Printf("Starting %s MCP server v%s", cfg.Server.Name, cfg.Server.Version)
	if err := server.Start(ctx); err != nil && err != context.Canceled {
		log.Fatal("Server error:", err)
	}

	log.Println("Server shut down gracefully")
}
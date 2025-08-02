package config

import (
	"os"
)

// Config holds all configuration for the application
type Config struct {
	Server ServerConfig
	Figma  FigmaConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Name    string
	Version string
}

// FigmaConfig holds Figma API related configuration
type FigmaConfig struct {
	APIKey string
	BaseURL string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	config := &Config{
		Server: ServerConfig{
			Name:    getEnv("SERVER_NAME", "Figma MCP Server"),
			Version: getEnv("SERVER_VERSION", "1.0.0"),
		},
		Figma: FigmaConfig{
			APIKey:  os.Getenv("FIGMA_API_KEY"),
			BaseURL: getEnv("FIGMA_BASE_URL", "https://api.figma.com/v1"),
		},
	}

	return config, nil
}

// getEnv returns the value of an environment variable or a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
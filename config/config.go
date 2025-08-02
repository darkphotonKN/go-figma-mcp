package config

import (
	"os"
)

// Config holds all configuration for the application
type Config struct {
	Server ServerConfig
	// TODO: Add other configuration sections as needed
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Name    string
	Version string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	config := &Config{
		Server: ServerConfig{
			Name:    getEnv("SERVER_NAME", "Figma MCP Server"),
			Version: getEnv("SERVER_VERSION", "1.0.0"),
		},
		// TODO: Load other configuration sections
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
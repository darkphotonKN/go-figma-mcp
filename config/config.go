package config

import (
	"fmt"
	"os"
)

type AppConfig struct {
	FigmaKey string
}

/**
* Loads app-wide configuration information.
**/
func LoadConfig() (*AppConfig, error) {
	figmaKey := getEnv("FIGMA_API_KEY", "")

	if figmaKey == "" {
		return nil, fmt.Errorf("Error when attempting to load Figma Key - key wasn't present.")
	}

	return &AppConfig{
		FigmaKey: figmaKey,
	}, nil
}

// getEnv returns the value of an environment variable or a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

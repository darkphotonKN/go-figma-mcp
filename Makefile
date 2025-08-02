# Build variables
BINARY_NAME=figma-mcp-server
BUILD_DIR=./bin
GO_FILES=$(shell find . -name '*.go' -type f)

# Go variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

.PHONY: all build clean run test deps tidy help

# Default target
all: clean build

# Build the binary
build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	@$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) cmd/main.go

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@$(GOCLEAN)
	@rm -rf $(BUILD_DIR)

# Run the application
run: build
	@echo "Running..."
	@$(BUILD_DIR)/$(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	@$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@$(GOTEST) -v -coverprofile=coverage.out ./...
	@$(GOCMD) tool cover -html=coverage.out

# Install dependencies
deps:
	@echo "Downloading dependencies..."
	@$(GOGET) -v ./...

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	@$(GOMOD) tidy

# Format Go code
fmt:
	@echo "Formatting code..."
	@$(GOCMD) fmt ./...

# Lint code (requires golangci-lint)
lint:
	@echo "Linting code..."
	@golangci-lint run

# Vet code
vet:
	@echo "Vetting code..."
	@$(GOCMD) vet ./...

# Build for multiple platforms
build-all: build-linux build-macos build-windows

build-linux:
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 cmd/main.go

build-macos:
	@echo "Building for macOS..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 cmd/main.go
	@GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 cmd/main.go

build-windows:
	@echo "Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe cmd/main.go

# Development mode with hot reload (requires air)
dev:
	@echo "Running in development mode..."
	@air

# Install development tools
install-tools:
	@echo "Installing development tools..."
	@$(GOGET) github.com/cosmtrek/air@latest
	@$(GOGET) github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Check if required environment variables are set
check-env:
	@echo "Checking environment variables..."
	@if [ -z "$$FIGMA_API_KEY" ]; then \
		echo "Error: FIGMA_API_KEY environment variable is not set"; \
		echo "Please copy .env.example to .env and set your Figma API key"; \
		exit 1; \
	fi
	@echo "Environment variables are properly set"

# Run with environment check
run-safe: check-env run

# Display help
help:
	@echo "Available targets:"
	@echo "  build         - Build the binary"
	@echo "  clean         - Clean build artifacts"
	@echo "  run           - Build and run the application"
	@echo "  run-safe      - Check environment and run the application"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  deps          - Download dependencies"
	@echo "  tidy          - Tidy dependencies"
	@echo "  fmt           - Format Go code"
	@echo "  lint          - Lint code (requires golangci-lint)"
	@echo "  vet           - Vet code"
	@echo "  build-all     - Build for multiple platforms"
	@echo "  dev           - Run in development mode with hot reload"
	@echo "  install-tools - Install development tools"
	@echo "  check-env     - Check if required environment variables are set"
	@echo "  help          - Display this help message"
# Go Figma MCP Server

A Model Context Protocol (MCP) server implementation for Figma integration.

## Project Structure

This project follows the Go MCP project template structure:

```
go-figma-mcp/
├── cmd/
│   └── main.go              # Application entry point
├── config/
│   └── config.go            # Configuration management
├── internal/
│   ├── constants/
│   │   └── api.go           # API constants
│   ├── figma/
│   │   ├── client.go        # External API client
│   │   ├── model.go         # Data models
│   │   ├── service.go       # Business logic
│   │   └── tools.go         # Request handlers
│   └── utils/
│       ├── errors.go        # Error handling utilities
│       └── validation.go    # Validation utilities
├── pkg/
│   └── mcp/
│       ├── capabilities.go  # MCP server capabilities
│       ├── server.go        # MCP server implementation
│       └── tools.go         # MCP tool framework
├── bin/                     # Compiled binaries (gitignored)
├── tmp/                     # Temporary files (gitignored)
├── .env.example             # Environment variables template
├── .gitignore
├── Makefile                 # Build and development commands
├── go.mod                   # Go module file
└── README.md                # This file
```

## Prerequisites

- Go 1.24.2 or later

## Getting Started

1. Clone the repository
2. Copy `.env.example` to `.env` and configure as needed
3. Install dependencies: `make deps`
4. Build the project: `make build`
5. Run the server: `make run`

## Development

- `make dev` - Run in development mode
- `make test` - Run tests
- `make fmt` - Format code
- `make lint` - Lint code

## TODO

- [ ] Implement MCP server functionality
- [ ] Add Figma API integration
- [ ] Define MCP tools and resources
- [ ] Add comprehensive tests
- [ ] Complete documentation
package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Server represents the MCP server instance
type Server struct {
	name        string
	version     string
	tools       map[string]Tool
	resources   map[string]Resource
	prompts     map[string]Prompt
	input       io.Reader
	output      io.Writer
	errOutput   io.Writer
	capabilities ServerCapabilities
}

// ServerConfig holds configuration for creating a new server
type ServerConfig struct {
	Name         string
	Version      string
	Input        io.Reader
	Output       io.Writer
	ErrOutput    io.Writer
	Capabilities ServerCapabilities
}

// NewServer creates a new MCP server instance
func NewServer(config ServerConfig) *Server {
	if config.Input == nil {
		config.Input = os.Stdin
	}
	if config.Output == nil {
		config.Output = os.Stdout
	}
	if config.ErrOutput == nil {
		config.ErrOutput = os.Stderr
	}

	return &Server{
		name:         config.Name,
		version:      config.Version,
		tools:        make(map[string]Tool),
		resources:    make(map[string]Resource),
		prompts:      make(map[string]Prompt),
		input:        config.Input,
		output:       config.Output,
		errOutput:    config.ErrOutput,
		capabilities: config.Capabilities,
	}
}

// RegisterTool registers a new tool with the server
func (s *Server) RegisterTool(tool Tool) error {
	if _, exists := s.tools[tool.Name]; exists {
		return fmt.Errorf("tool %s already registered", tool.Name)
	}
	s.tools[tool.Name] = tool
	return nil
}

// RegisterResource registers a new resource with the server
func (s *Server) RegisterResource(resource Resource) error {
	if _, exists := s.resources[resource.URI]; exists {
		return fmt.Errorf("resource %s already registered", resource.URI)
	}
	s.resources[resource.URI] = resource
	return nil
}

// RegisterPrompt registers a new prompt with the server
func (s *Server) RegisterPrompt(prompt Prompt) error {
	if _, exists := s.prompts[prompt.Name]; exists {
		return fmt.Errorf("prompt %s already registered", prompt.Name)
	}
	s.prompts[prompt.Name] = prompt
	return nil
}

// Start begins the MCP server and handles incoming requests
func (s *Server) Start(ctx context.Context) error {
	log.Printf("Starting MCP server: %s v%s", s.name, s.version)

	decoder := json.NewDecoder(s.input)
	encoder := json.NewEncoder(s.output)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			var msg Message
			if err := decoder.Decode(&msg); err != nil {
				if err == io.EOF {
					return nil
				}
				return fmt.Errorf("failed to decode message: %w", err)
			}

			response, err := s.handleMessage(ctx, msg)
			if err != nil {
				s.sendError(encoder, msg.ID, err)
				continue
			}

			if response != nil {
				if err := encoder.Encode(response); err != nil {
					return fmt.Errorf("failed to encode response: %w", err)
				}
			}
		}
	}
}

// handleMessage processes incoming MCP messages
func (s *Server) handleMessage(ctx context.Context, msg Message) (*Message, error) {
	switch msg.Method {
	case "initialize":
		return s.handleInitialize(msg)
	case "tools/list":
		return s.handleToolsList(msg)
	case "tools/call":
		return s.handleToolCall(ctx, msg)
	case "resources/list":
		return s.handleResourcesList(msg)
	case "resources/read":
		return s.handleResourceRead(ctx, msg)
	case "prompts/list":
		return s.handlePromptsList(msg)
	case "prompts/get":
		return s.handlePromptGet(msg)
	case "completion/complete":
		return s.handleCompletion(ctx, msg)
	default:
		return nil, fmt.Errorf("unknown method: %s", msg.Method)
	}
}

// handleInitialize handles the initialize request
func (s *Server) handleInitialize(msg Message) (*Message, error) {
	result := InitializeResult{
		ProtocolVersion: "2024-11-05",
		ServerInfo: ServerInfo{
			Name:    s.name,
			Version: s.version,
		},
		Capabilities: s.capabilities,
	}

	return &Message{
		JSONRPC: "2.0",
		ID:      msg.ID,
		Result:  result,
	}, nil
}

// handleToolsList returns the list of available tools
func (s *Server) handleToolsList(msg Message) (*Message, error) {
	tools := make([]Tool, 0, len(s.tools))
	for _, tool := range s.tools {
		tools = append(tools, tool)
	}

	return &Message{
		JSONRPC: "2.0",
		ID:      msg.ID,
		Result: ToolsListResult{
			Tools: tools,
		},
	}, nil
}

// handleToolCall executes a tool and returns the result
func (s *Server) handleToolCall(ctx context.Context, msg Message) (*Message, error) {
	var params ToolCallParams
	if err := json.Unmarshal(msg.Params, &params); err != nil {
		return nil, fmt.Errorf("invalid tool call params: %w", err)
	}

	tool, exists := s.tools[params.Name]
	if !exists {
		return nil, fmt.Errorf("tool not found: %s", params.Name)
	}

	result, err := tool.Handler(ctx, params.Arguments)
	if err != nil {
		return nil, fmt.Errorf("tool execution failed: %w", err)
	}

	return &Message{
		JSONRPC: "2.0",
		ID:      msg.ID,
		Result: ToolCallResult{
			Content: []Content{
				{
					Type: "text",
					Text: result,
				},
			},
		},
	}, nil
}

// handleResourcesList returns the list of available resources
func (s *Server) handleResourcesList(msg Message) (*Message, error) {
	resources := make([]Resource, 0, len(s.resources))
	for _, resource := range s.resources {
		resources = append(resources, resource)
	}

	return &Message{
		JSONRPC: "2.0",
		ID:      msg.ID,
		Result: ResourcesListResult{
			Resources: resources,
		},
	}, nil
}

// handleResourceRead reads a resource and returns its content
func (s *Server) handleResourceRead(ctx context.Context, msg Message) (*Message, error) {
	var params ResourceReadParams
	if err := json.Unmarshal(msg.Params, &params); err != nil {
		return nil, fmt.Errorf("invalid resource read params: %w", err)
	}

	resource, exists := s.resources[params.URI]
	if !exists {
		return nil, fmt.Errorf("resource not found: %s", params.URI)
	}

	content, err := resource.Handler(ctx, params.URI)
	if err != nil {
		return nil, fmt.Errorf("resource read failed: %w", err)
	}

	return &Message{
		JSONRPC: "2.0",
		ID:      msg.ID,
		Result: ResourceReadResult{
			Contents: []ResourceContent{
				{
					URI:      params.URI,
					MimeType: resource.MimeType,
					Text:     content,
				},
			},
		},
	}, nil
}

// handlePromptsList returns the list of available prompts
func (s *Server) handlePromptsList(msg Message) (*Message, error) {
	prompts := make([]Prompt, 0, len(s.prompts))
	for _, prompt := range s.prompts {
		prompts = append(prompts, prompt)
	}

	return &Message{
		JSONRPC: "2.0",
		ID:      msg.ID,
		Result: PromptsListResult{
			Prompts: prompts,
		},
	}, nil
}

// handlePromptGet retrieves a specific prompt
func (s *Server) handlePromptGet(msg Message) (*Message, error) {
	var params PromptGetParams
	if err := json.Unmarshal(msg.Params, &params); err != nil {
		return nil, fmt.Errorf("invalid prompt get params: %w", err)
	}

	prompt, exists := s.prompts[params.Name]
	if !exists {
		return nil, fmt.Errorf("prompt not found: %s", params.Name)
	}

	// Generate prompt messages based on arguments
	messages, err := prompt.Handler(params.Arguments)
	if err != nil {
		return nil, fmt.Errorf("prompt generation failed: %w", err)
	}

	return &Message{
		JSONRPC: "2.0",
		ID:      msg.ID,
		Result: PromptGetResult{
			Messages: messages,
		},
	}, nil
}

// handleCompletion handles completion requests
func (s *Server) handleCompletion(ctx context.Context, msg Message) (*Message, error) {
	if !s.capabilities.Completion.Provider {
		return nil, fmt.Errorf("completion not supported")
	}

	var params CompletionParams
	if err := json.Unmarshal(msg.Params, &params); err != nil {
		return nil, fmt.Errorf("invalid completion params: %w", err)
	}

	// TODO: Implement completion logic based on your needs
	// This is a placeholder implementation
	return &Message{
		JSONRPC: "2.0",
		ID:      msg.ID,
		Result: CompletionResult{
			Completion: CompletionOption{
				Values:  []string{},
				Total:   0,
				HasMore: false,
			},
		},
	}, nil
}

// sendError sends an error response
func (s *Server) sendError(encoder *json.Encoder, id json.RawMessage, err error) {
	errResp := &Message{
		JSONRPC: "2.0",
		ID:      id,
		Error: &Error{
			Code:    -32603, // Internal error
			Message: err.Error(),
		},
	}
	encoder.Encode(errResp)
}

// Message represents a JSON-RPC message
type Message struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Result  interface{}     `json:"result,omitempty"`
	Error   *Error          `json:"error,omitempty"`
}

// Error represents a JSON-RPC error
type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// InitializeResult represents the response to an initialize request
type InitializeResult struct {
	ProtocolVersion string             `json:"protocolVersion"`
	ServerInfo      ServerInfo         `json:"serverInfo"`
	Capabilities    ServerCapabilities `json:"capabilities"`
}

// ServerInfo contains information about the server
type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ToolsListResult represents the response to a tools/list request
type ToolsListResult struct {
	Tools []Tool `json:"tools"`
}

// ToolCallParams represents parameters for a tool call
type ToolCallParams struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

// ToolCallResult represents the result of a tool call
type ToolCallResult struct {
	Content []Content `json:"content"`
}

// Content represents content in a response
type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// ResourcesListResult represents the response to a resources/list request
type ResourcesListResult struct {
	Resources []Resource `json:"resources"`
}

// ResourceReadParams represents parameters for reading a resource
type ResourceReadParams struct {
	URI string `json:"uri"`
}

// ResourceReadResult represents the result of reading a resource
type ResourceReadResult struct {
	Contents []ResourceContent `json:"contents"`
}

// ResourceContent represents the content of a resource
type ResourceContent struct {
	URI      string `json:"uri"`
	MimeType string `json:"mimeType"`
	Text     string `json:"text,omitempty"`
	Blob     string `json:"blob,omitempty"`
}

// PromptsListResult represents the response to a prompts/list request
type PromptsListResult struct {
	Prompts []Prompt `json:"prompts"`
}

// PromptGetParams represents parameters for getting a prompt
type PromptGetParams struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments,omitempty"`
}

// PromptGetResult represents the result of getting a prompt
type PromptGetResult struct {
	Messages []PromptMessage `json:"messages"`
}

// PromptMessage represents a message in a prompt
type PromptMessage struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

// CompletionParams represents parameters for completion
type CompletionParams struct {
	Ref       CompletionRef                  `json:"ref"`
	Argument  CompletionArgument             `json:"argument"`
}

// CompletionRef represents a reference for completion
type CompletionRef struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// CompletionArgument represents an argument for completion
type CompletionArgument struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// CompletionResult represents the result of a completion request
type CompletionResult struct {
	Completion CompletionOption `json:"completion"`
}

// CompletionOption represents completion options
type CompletionOption struct {
	Values  []string `json:"values"`
	Total   int      `json:"total"`
	HasMore bool     `json:"hasMore"`
}
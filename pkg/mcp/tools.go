package mcp

import (
	"context"
	"encoding/json"
)

// Tool represents an MCP tool that can be called by clients
type Tool struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	InputSchema json.RawMessage `json:"inputSchema"`
	Handler     ToolHandler     `json:"-"`
}

// ToolHandler is a function that handles tool execution
type ToolHandler func(ctx context.Context, arguments map[string]interface{}) (string, error)

// Resource represents an MCP resource that can be read
type Resource struct {
	URI         string          `json:"uri"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	MimeType    string          `json:"mimeType,omitempty"`
	Handler     ResourceHandler `json:"-"`
}

// ResourceHandler is a function that handles resource reading
type ResourceHandler func(ctx context.Context, uri string) (string, error)

// Prompt represents an MCP prompt template
type Prompt struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Arguments   []PromptArgument       `json:"arguments,omitempty"`
	Handler     PromptHandler          `json:"-"`
}

// PromptArgument represents an argument for a prompt
type PromptArgument struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
}

// PromptHandler is a function that generates prompt messages
type PromptHandler func(arguments map[string]interface{}) ([]PromptMessage, error)

// ToolBuilder helps construct tools with proper JSON schemas
type ToolBuilder struct {
	name        string
	description string
	schema      map[string]interface{}
	handler     ToolHandler
}

// NewToolBuilder creates a new tool builder
func NewToolBuilder(name, description string) *ToolBuilder {
	return &ToolBuilder{
		name:        name,
		description: description,
		schema: map[string]interface{}{
			"type": "object",
			"properties": make(map[string]interface{}),
			"required": []string{},
		},
	}
}

// AddProperty adds a property to the tool's input schema
func (tb *ToolBuilder) AddProperty(name, propType, description string, required bool) *ToolBuilder {
	properties := tb.schema["properties"].(map[string]interface{})
	properties[name] = map[string]interface{}{
		"type":        propType,
		"description": description,
	}

	if required {
		requiredList := tb.schema["required"].([]string)
		tb.schema["required"] = append(requiredList, name)
	}

	return tb
}

// AddStringProperty adds a string property to the tool's input schema
func (tb *ToolBuilder) AddStringProperty(name, description string, required bool) *ToolBuilder {
	return tb.AddProperty(name, "string", description, required)
}

// AddNumberProperty adds a number property to the tool's input schema
func (tb *ToolBuilder) AddNumberProperty(name, description string, required bool) *ToolBuilder {
	return tb.AddProperty(name, "number", description, required)
}

// AddBooleanProperty adds a boolean property to the tool's input schema
func (tb *ToolBuilder) AddBooleanProperty(name, description string, required bool) *ToolBuilder {
	return tb.AddProperty(name, "boolean", description, required)
}

// AddArrayProperty adds an array property to the tool's input schema
func (tb *ToolBuilder) AddArrayProperty(name, itemType, description string, required bool) *ToolBuilder {
	properties := tb.schema["properties"].(map[string]interface{})
	properties[name] = map[string]interface{}{
		"type":        "array",
		"description": description,
		"items": map[string]interface{}{
			"type": itemType,
		},
	}

	if required {
		requiredList := tb.schema["required"].([]string)
		tb.schema["required"] = append(requiredList, name)
	}

	return tb
}

// AddObjectProperty adds an object property to the tool's input schema
func (tb *ToolBuilder) AddObjectProperty(name, description string, properties map[string]interface{}, required bool) *ToolBuilder {
	props := tb.schema["properties"].(map[string]interface{})
	props[name] = map[string]interface{}{
		"type":        "object",
		"description": description,
		"properties":  properties,
	}

	if required {
		requiredList := tb.schema["required"].([]string)
		tb.schema["required"] = append(requiredList, name)
	}

	return tb
}

// SetHandler sets the handler function for the tool
func (tb *ToolBuilder) SetHandler(handler ToolHandler) *ToolBuilder {
	tb.handler = handler
	return tb
}

// Build creates the final Tool instance
func (tb *ToolBuilder) Build() (Tool, error) {
	schemaJSON, err := json.Marshal(tb.schema)
	if err != nil {
		return Tool{}, err
	}

	return Tool{
		Name:        tb.name,
		Description: tb.description,
		InputSchema: schemaJSON,
		Handler:     tb.handler,
	}, nil
}

// ResourceBuilder helps construct resources
type ResourceBuilder struct {
	uri         string
	name        string
	description string
	mimeType    string
	handler     ResourceHandler
}

// NewResourceBuilder creates a new resource builder
func NewResourceBuilder(uri, name string) *ResourceBuilder {
	return &ResourceBuilder{
		uri:      uri,
		name:     name,
		mimeType: "text/plain",
	}
}

// SetDescription sets the resource description
func (rb *ResourceBuilder) SetDescription(description string) *ResourceBuilder {
	rb.description = description
	return rb
}

// SetMimeType sets the resource MIME type
func (rb *ResourceBuilder) SetMimeType(mimeType string) *ResourceBuilder {
	rb.mimeType = mimeType
	return rb
}

// SetHandler sets the handler function for the resource
func (rb *ResourceBuilder) SetHandler(handler ResourceHandler) *ResourceBuilder {
	rb.handler = handler
	return rb
}

// Build creates the final Resource instance
func (rb *ResourceBuilder) Build() Resource {
	return Resource{
		URI:         rb.uri,
		Name:        rb.name,
		Description: rb.description,
		MimeType:    rb.mimeType,
		Handler:     rb.handler,
	}
}

// PromptBuilder helps construct prompts
type PromptBuilder struct {
	name        string
	description string
	arguments   []PromptArgument
	handler     PromptHandler
}

// NewPromptBuilder creates a new prompt builder
func NewPromptBuilder(name string) *PromptBuilder {
	return &PromptBuilder{
		name:      name,
		arguments: []PromptArgument{},
	}
}

// SetDescription sets the prompt description
func (pb *PromptBuilder) SetDescription(description string) *PromptBuilder {
	pb.description = description
	return pb
}

// AddArgument adds an argument to the prompt
func (pb *PromptBuilder) AddArgument(name, description string, required bool) *PromptBuilder {
	pb.arguments = append(pb.arguments, PromptArgument{
		Name:        name,
		Description: description,
		Required:    required,
	})
	return pb
}

// SetHandler sets the handler function for the prompt
func (pb *PromptBuilder) SetHandler(handler PromptHandler) *PromptBuilder {
	pb.handler = handler
	return pb
}

// Build creates the final Prompt instance
func (pb *PromptBuilder) Build() Prompt {
	return Prompt{
		Name:        pb.name,
		Description: pb.description,
		Arguments:   pb.arguments,
		Handler:     pb.handler,
	}
}

// Common validation helpers

// ValidateRequiredString validates that a required string argument is present and non-empty
func ValidateRequiredString(args map[string]interface{}, key string) (string, error) {
	value, exists := args[key]
	if !exists {
		return "", &ValidationError{Field: key, Message: "required field missing"}
	}

	str, ok := value.(string)
	if !ok {
		return "", &ValidationError{Field: key, Message: "must be a string"}
	}

	if str == "" {
		return "", &ValidationError{Field: key, Message: "cannot be empty"}
	}

	return str, nil
}

// ValidateOptionalString validates an optional string argument
func ValidateOptionalString(args map[string]interface{}, key string, defaultValue string) string {
	value, exists := args[key]
	if !exists {
		return defaultValue
	}

	str, ok := value.(string)
	if !ok {
		return defaultValue
	}

	return str
}

// ValidateRequiredNumber validates that a required number argument is present
func ValidateRequiredNumber(args map[string]interface{}, key string) (float64, error) {
	value, exists := args[key]
	if !exists {
		return 0, &ValidationError{Field: key, Message: "required field missing"}
	}

	num, ok := value.(float64)
	if !ok {
		// Try to convert from int
		if intVal, ok := value.(int); ok {
			return float64(intVal), nil
		}
		return 0, &ValidationError{Field: key, Message: "must be a number"}
	}

	return num, nil
}

// ValidateOptionalNumber validates an optional number argument
func ValidateOptionalNumber(args map[string]interface{}, key string, defaultValue float64) float64 {
	value, exists := args[key]
	if !exists {
		return defaultValue
	}

	num, ok := value.(float64)
	if !ok {
		// Try to convert from int
		if intVal, ok := value.(int); ok {
			return float64(intVal)
		}
		return defaultValue
	}

	return num
}

// ValidateRequiredBool validates that a required boolean argument is present
func ValidateRequiredBool(args map[string]interface{}, key string) (bool, error) {
	value, exists := args[key]
	if !exists {
		return false, &ValidationError{Field: key, Message: "required field missing"}
	}

	boolVal, ok := value.(bool)
	if !ok {
		return false, &ValidationError{Field: key, Message: "must be a boolean"}
	}

	return boolVal, nil
}

// ValidateOptionalBool validates an optional boolean argument
func ValidateOptionalBool(args map[string]interface{}, key string, defaultValue bool) bool {
	value, exists := args[key]
	if !exists {
		return defaultValue
	}

	boolVal, ok := value.(bool)
	if !ok {
		return defaultValue
	}

	return boolVal
}

// ValidationError represents a validation error for tool arguments
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return "validation error for field '" + e.Field + "': " + e.Message
}
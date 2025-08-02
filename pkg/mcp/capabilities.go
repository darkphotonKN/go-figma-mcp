package mcp

// ServerCapabilities represents the capabilities of an MCP server
type ServerCapabilities struct {
	// Tools capability indicates if the server supports tool calls
	Tools *ToolsCapability `json:"tools,omitempty"`

	// Resources capability indicates if the server provides resources
	Resources *ResourcesCapability `json:"resources,omitempty"`

	// Prompts capability indicates if the server supports prompts
	Prompts *PromptsCapability `json:"prompts,omitempty"`

	// Logging capability indicates if the server supports logging
	Logging *LoggingCapability `json:"logging,omitempty"`

	// Completion capability indicates if the server supports completion
	Completion *CompletionCapability `json:"completion,omitempty"`
}

// ToolsCapability represents the tools capability
type ToolsCapability struct {
	// Provider indicates if the server can provide tools
	Provider bool `json:"provider"`
}

// ResourcesCapability represents the resources capability
type ResourcesCapability struct {
	// Provider indicates if the server can provide resources
	Provider bool `json:"provider"`

	// Subscribe indicates if the server supports resource subscriptions
	Subscribe bool `json:"subscribe,omitempty"`
}

// PromptsCapability represents the prompts capability
type PromptsCapability struct {
	// Provider indicates if the server can provide prompts
	Provider bool `json:"provider"`
}

// LoggingCapability represents the logging capability
type LoggingCapability struct {
	// Provider indicates if the server can provide logging
	Provider bool `json:"provider"`
}

// CompletionCapability represents the completion capability
type CompletionCapability struct {
	// Provider indicates if the server can provide completion
	Provider bool `json:"provider"`
}

// DefaultCapabilities returns a default set of capabilities with all features disabled
func DefaultCapabilities() ServerCapabilities {
	return ServerCapabilities{
		Tools: &ToolsCapability{
			Provider: false,
		},
		Resources: &ResourcesCapability{
			Provider:  false,
			Subscribe: false,
		},
		Prompts: &PromptsCapability{
			Provider: false,
		},
		Logging: &LoggingCapability{
			Provider: false,
		},
		Completion: &CompletionCapability{
			Provider: false,
		},
	}
}

// WithTools returns capabilities with tools enabled
func WithTools() ServerCapabilities {
	caps := DefaultCapabilities()
	caps.Tools.Provider = true
	return caps
}

// WithResources returns capabilities with resources enabled
func WithResources(subscribe bool) ServerCapabilities {
	caps := DefaultCapabilities()
	caps.Resources.Provider = true
	caps.Resources.Subscribe = subscribe
	return caps
}

// WithPrompts returns capabilities with prompts enabled
func WithPrompts() ServerCapabilities {
	caps := DefaultCapabilities()
	caps.Prompts.Provider = true
	return caps
}

// WithLogging returns capabilities with logging enabled
func WithLogging() ServerCapabilities {
	caps := DefaultCapabilities()
	caps.Logging.Provider = true
	return caps
}

// WithCompletion returns capabilities with completion enabled
func WithCompletion() ServerCapabilities {
	caps := DefaultCapabilities()
	caps.Completion.Provider = true
	return caps
}

// AllCapabilities returns capabilities with all features enabled
func AllCapabilities() ServerCapabilities {
	return ServerCapabilities{
		Tools: &ToolsCapability{
			Provider: true,
		},
		Resources: &ResourcesCapability{
			Provider:  true,
			Subscribe: true,
		},
		Prompts: &PromptsCapability{
			Provider: true,
		},
		Logging: &LoggingCapability{
			Provider: true,
		},
		Completion: &CompletionCapability{
			Provider: true,
		},
	}
}

// CapabilitiesBuilder helps construct custom server capabilities
type CapabilitiesBuilder struct {
	capabilities ServerCapabilities
}

// NewCapabilitiesBuilder creates a new capabilities builder
func NewCapabilitiesBuilder() *CapabilitiesBuilder {
	return &CapabilitiesBuilder{
		capabilities: DefaultCapabilities(),
	}
}

// EnableTools enables the tools capability
func (cb *CapabilitiesBuilder) EnableTools() *CapabilitiesBuilder {
	cb.capabilities.Tools.Provider = true
	return cb
}

// EnableResources enables the resources capability
func (cb *CapabilitiesBuilder) EnableResources(withSubscribe bool) *CapabilitiesBuilder {
	cb.capabilities.Resources.Provider = true
	cb.capabilities.Resources.Subscribe = withSubscribe
	return cb
}

// EnablePrompts enables the prompts capability
func (cb *CapabilitiesBuilder) EnablePrompts() *CapabilitiesBuilder {
	cb.capabilities.Prompts.Provider = true
	return cb
}

// EnableLogging enables the logging capability
func (cb *CapabilitiesBuilder) EnableLogging() *CapabilitiesBuilder {
	cb.capabilities.Logging.Provider = true
	return cb
}

// EnableCompletion enables the completion capability
func (cb *CapabilitiesBuilder) EnableCompletion() *CapabilitiesBuilder {
	cb.capabilities.Completion.Provider = true
	return cb
}

// Build returns the constructed capabilities
func (cb *CapabilitiesBuilder) Build() ServerCapabilities {
	return cb.capabilities
}

// Merge combines two sets of capabilities, with the second set taking precedence
func (c ServerCapabilities) Merge(other ServerCapabilities) ServerCapabilities {
	result := c

	if other.Tools != nil {
		if result.Tools == nil {
			result.Tools = &ToolsCapability{}
		}
		result.Tools.Provider = other.Tools.Provider
	}

	if other.Resources != nil {
		if result.Resources == nil {
			result.Resources = &ResourcesCapability{}
		}
		result.Resources.Provider = other.Resources.Provider
		result.Resources.Subscribe = other.Resources.Subscribe
	}

	if other.Prompts != nil {
		if result.Prompts == nil {
			result.Prompts = &PromptsCapability{}
		}
		result.Prompts.Provider = other.Prompts.Provider
	}

	if other.Logging != nil {
		if result.Logging == nil {
			result.Logging = &LoggingCapability{}
		}
		result.Logging.Provider = other.Logging.Provider
	}

	if other.Completion != nil {
		if result.Completion == nil {
			result.Completion = &CompletionCapability{}
		}
		result.Completion.Provider = other.Completion.Provider
	}

	return result
}

// HasAnyCapability returns true if any capability is enabled
func (c ServerCapabilities) HasAnyCapability() bool {
	return (c.Tools != nil && c.Tools.Provider) ||
		(c.Resources != nil && c.Resources.Provider) ||
		(c.Prompts != nil && c.Prompts.Provider) ||
		(c.Logging != nil && c.Logging.Provider) ||
		(c.Completion != nil && c.Completion.Provider)
}

// ValidateCapabilities ensures the capabilities are properly configured
func (c ServerCapabilities) ValidateCapabilities() error {
	// Add any validation logic here if needed
	// For example, checking for conflicting capabilities or required combinations
	return nil
}
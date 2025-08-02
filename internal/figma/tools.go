package figma

// Handler handles requests for the domain
type Handler struct {
	service Service
}

// NewHandler creates a new handler instance
func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// TODO: Add MCP tool registration and handler methods
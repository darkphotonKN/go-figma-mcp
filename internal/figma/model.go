package figma

// Entity represents the main domain entity
type Entity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateEntityRequest represents the request to create an entity
type CreateEntityRequest struct {
	Name string `json:"name" binding:"required"`
}

// UpdateEntityRequest represents the request to update an entity
type UpdateEntityRequest struct {
	Name string `json:"name"`
}

// TODO: Add Figma-specific models as needed
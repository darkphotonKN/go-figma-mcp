package figma

import (
	"context"
)

// Service interface defines the business logic methods
type Service interface {
	GetFileInfo(ctx context.Context, fileID string) error
	// TODO: Add other business logic methods as needed
}

// service implements the Service interface
type service struct {
	client *Client
}

// NewService creates a new service instance
func NewService(client *Client) Service {
	return &service{
		client: client,
	}
}

func (s *service) GetFileInfo(ctx context.Context, fileID string) error {
	return s.client.GetFileInfo(fileID)
}
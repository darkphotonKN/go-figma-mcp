package figma

import (
	"context"
)

type Service interface {
	GetFileInfo(ctx context.Context, fileID string) error
}

type service struct {
	client *Client
}

func NewService(client *Client) Service {
	return &service{
		client: client,
	}
}

func (s *service) GetFileInfo(ctx context.Context, fileID string) error {
	return s.client.GetFileInfo(fileID)
}

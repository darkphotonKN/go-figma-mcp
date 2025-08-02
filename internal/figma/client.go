package figma

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		baseURL:    "https://api.figma.com/v1",
		apiKey:     apiKey,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) GetFileInfo(fileID string) error {
	// TODO: Implement actual Figma API call
	fmt.Printf("Would fetch file info for: %s\n", fileID)
	return nil
}

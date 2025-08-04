package figma

import (
	"fmt"
	"io"
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
	err := c.fetchFigmaFile("C1saDjsNsINCe5nj73eJXL")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) fetchFigmaFile(fileKey string) error {
	url := fmt.Sprintf("https://api.figma.com/v1/files/%s", fileKey)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("\n\nFigma File Response: %s\n\n", string(body))

	return nil
}

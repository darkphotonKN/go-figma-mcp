package figma

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Client represents a Figma API client
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new Figma API client
func NewClient(apiKey, baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetFile retrieves a Figma file by its key
func (c *Client) GetFile(req GetFileRequest) (*FileResponse, error) {
	endpoint := fmt.Sprintf("/files/%s", req.FileKey)
	
	params := url.Values{}
	if req.Version != "" {
		params.Add("version", req.Version)
	}
	if req.IDs != "" {
		params.Add("ids", req.IDs)
	}
	if req.Depth > 0 {
		params.Add("depth", fmt.Sprintf("%d", req.Depth))
	}

	var response FileResponse
	err := c.makeRequest("GET", endpoint, params, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get file: %w", err)
	}

	return &response, nil
}

// GetImages retrieves rendered images for specified nodes
func (c *Client) GetImages(req GetImageRequest) (*ImageResponse, error) {
	endpoint := fmt.Sprintf("/images/%s", req.FileKey)
	
	params := url.Values{}
	params.Add("ids", req.IDs)
	if req.Scale != "" {
		params.Add("scale", req.Scale)
	}
	if req.Format != "" {
		params.Add("format", req.Format)
	}
	if req.UseAbsoluteBounds {
		params.Add("use_absolute_bounds", "true")
	}

	var response ImageResponse
	err := c.makeRequest("GET", endpoint, params, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get images: %w", err)
	}

	if response.Err != nil {
		return nil, fmt.Errorf("figma API error: %s", *response.Err)
	}

	return &response, nil
}

// GetComments retrieves comments for a file
func (c *Client) GetComments(req CommentRequest) (*CommentsResponse, error) {
	endpoint := fmt.Sprintf("/files/%s/comments", req.FileKey)
	
	var response CommentsResponse
	err := c.makeRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments: %w", err)
	}

	return &response, nil
}

// makeRequest makes an HTTP request to the Figma API
func (c *Client) makeRequest(method, endpoint string, params url.Values, result interface{}) error {
	reqURL := c.baseURL + endpoint
	if params != nil && len(params) > 0 {
		reqURL += "?" + params.Encode()
	}

	req, err := http.NewRequest(method, reqURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Figma-Token", c.apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}
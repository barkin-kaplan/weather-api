package custom_http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	headers    map[string]string
}

// NewClient creates a new HTTP client with optional base URL and timeout
func NewClient(baseURL string, timeout time.Duration, defaultHeaders map[string]string) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: timeout},
		baseURL:    baseURL,
		headers:    defaultHeaders,
	}
}

// Get sends a GET request and decodes the response JSON into the result
func (c *Client) Get(ctx context.Context, path string, result interface{}, headers map[string]string) error {
	return c.doRequest(ctx, http.MethodGet, path, nil, result, headers)
}

// Post sends a POST request with JSON body and decodes the response JSON into the result
func (c *Client) Post(ctx context.Context, path string, body interface{}, result interface{}, headers map[string]string) error {
	return c.doRequest(ctx, http.MethodPost, path, body, result, headers)
}

func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}, result interface{}, headers map[string]string) error {
	fullURL := c.baseURL + path

	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewBuffer(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set default and request-specific headers
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}

	return nil
}

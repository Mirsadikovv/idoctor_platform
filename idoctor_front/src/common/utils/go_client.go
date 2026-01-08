package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ClientConfig struct {
	BaseURL  string
	BasePath string // optional
	APIKey   string // optional
	Name     string // identifier (EGOV, RUGOV, etc.)
}

type HTTPClient struct {
	BaseURL  string
	BasePath string // optional: e.g. "egov_dmi"
	APIKey   string
	client   *http.Client
	Name     string
}

func NewHTTPClient(cfg ClientConfig) *HTTPClient {
	if cfg.BaseURL == "" {
		panic(fmt.Sprintf("[%s] BaseURL is required", cfg.Name))
	}

	return &HTTPClient{
		BaseURL:  cfg.BaseURL,
		BasePath: cfg.BasePath,
		APIKey:   cfg.APIKey,
		client:   &http.Client{},
		Name:     cfg.Name,
	}
}

func (c *HTTPClient) DoRequest(method, path string, query map[string]any, body map[string]any) ([]byte, int, error) {
	var fullURL string
	path = url.PathEscape(path)
	if c.BasePath != "" {
		fullURL = fmt.Sprintf("%s/%s/%s", c.BaseURL, c.BasePath, path)
	} else {
		fullURL = fmt.Sprintf("%s/%s", c.BaseURL, path)
	}

	fullURL = strings.TrimSpace(fullURL)
	u, err := url.Parse(fullURL)
	if err != nil {
		return nil, 0, err
	}

	// Add query params if provided

	if query != nil {
		q := u.Query()
		for k, v := range query {
			switch vv := v.(type) {
			case string:
				q.Set(k, vv)
			case fmt.Stringer: // any type that implements String() string
				q.Set(k, vv.String())
			default:
				q.Set(k, fmt.Sprintf("%v", vv)) // fallback
			}
		}
		u.RawQuery = q.Encode()
	}

	// Encode body if present
	var buf io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, 0, err
		}
		buf = bytes.NewReader(b)
	}

	// Build request
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.APIKey != "" {
		req.Header.Set("api-key", c.APIKey)
	}

	// Send
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)

	return res, resp.StatusCode, err
}

func (c *HTTPClient) Get(path string, query map[string]any) ([]byte, int, error) {
	return c.DoRequest(http.MethodGet, path, query, nil)
}

// POST request
func (c *HTTPClient) Post(path string, query map[string]any, body map[string]any) ([]byte, int, error) {
	return c.DoRequest(http.MethodPost, path, query, body)
}

// PUT request
func (c *HTTPClient) Put(path string, query map[string]any, body map[string]any) ([]byte, int, error) {
	return c.DoRequest(http.MethodPut, path, query, body)
}

// PATCH request
func (c *HTTPClient) Patch(path string, query map[string]any, body map[string]any) ([]byte, int, error) {
	return c.DoRequest(http.MethodPatch, path, query, body)
}

// DELETE request
func (c *HTTPClient) Delete(path string, query map[string]any) ([]byte, int, error) {
	return c.DoRequest(http.MethodDelete, path, query, nil)
}

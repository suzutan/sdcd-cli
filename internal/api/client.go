package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/suzutan/sdcd-cli/internal/config"
)

// Client is an authenticated HTTP client for Screwdriver.cd API.
type Client struct {
	baseURL    string
	apiToken   string
	jwt        string
	httpClient *http.Client
}

// NewClient creates a new Client from a config Context.
func NewClient(ctx *config.Context) *Client {
	return &Client{
		baseURL:  ctx.APIURL,
		apiToken: ctx.Token,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// authenticate exchanges the raw API token for a JWT (cached in memory).
func (c *Client) authenticate() error {
	if c.jwt != "" {
		return nil
	}
	u := fmt.Sprintf("%s/v4/auth/token?api_token=%s", c.baseURL, url.QueryEscape(c.apiToken))
	resp, err := c.httpClient.Get(u) //nolint:noctx
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("auth: HTTP %d: %s", resp.StatusCode, string(body))
	}
	var result struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("auth decode: %w", err)
	}
	c.jwt = result.Token
	return nil
}

// do performs an authenticated HTTP request.
// body is JSON-encoded if non-nil. result is JSON-decoded from response if non-nil.
func (c *Client) do(method, path string, body, result interface{}) error {
	if err := c.authenticate(); err != nil {
		return err
	}

	var reqBody io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request: %w", err)
		}
		reqBody = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, c.baseURL+path, reqBody)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.jwt)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("http do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(respBody))
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}
	return nil
}

// doWithHeaders performs an authenticated HTTP request and returns response headers.
func (c *Client) doWithHeaders(method, path string, body, result interface{}) (http.Header, error) {
	if err := c.authenticate(); err != nil {
		return nil, err
	}

	var reqBody io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request: %w", err)
		}
		reqBody = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, c.baseURL+path, reqBody)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.jwt)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(respBody))
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return nil, fmt.Errorf("decode response: %w", err)
		}
	}
	return resp.Header, nil
}

func (c *Client) get(path string, result interface{}) error {
	return c.do(http.MethodGet, path, nil, result)
}

func (c *Client) post(path string, body, result interface{}) error {
	return c.do(http.MethodPost, path, body, result)
}

func (c *Client) put(path string, body, result interface{}) error {
	return c.do(http.MethodPut, path, body, result)
}

func (c *Client) delete(path string) error {
	return c.do(http.MethodDelete, path, nil, nil)
}

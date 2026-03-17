package valence

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Client is the main entry point for the Valence API.
type Client struct {
	baseURL    string
	auth       Authenticator
	httpClient *http.Client
	rateLimit  *rateLimitState
	stats      *Stats

	LPVersion  string
	LEVersion  string
	BASVersion string
}

// Config holds the configuration for creating a new Client.
type Config struct {
	// BaseURL is the root URL of the Brightspace instance, e.g. "https://learn.example.com"
	BaseURL    string
	Auth       Authenticator
	LPVersion  string
	LEVersion  string
	BASVersion string
	// HTTPClient allows supplying a custom http.Client (optional).
	HTTPClient *http.Client
}

// New creates a new Valence API client.
func New(cfg Config) *Client {
	hc := cfg.HTTPClient
	if hc == nil {
		hc = &http.Client{Timeout: 60 * time.Second}
	}
	return &Client{
		baseURL:    cfg.BaseURL,
		auth:       cfg.Auth,
		httpClient: hc,
		rateLimit:  &rateLimitState{},
		stats:      &Stats{},
		LPVersion:  cfg.LPVersion,
		LEVersion:  cfg.LEVersion,
		BASVersion: cfg.BASVersion,
	}
}

// Stats returns the API usage statistics.
func (c *Client) Stats() *Stats {
	return c.stats
}

// lpPath builds a path for an LP API route.
func (c *Client) lpPath(format string, args ...any) string {
	return fmt.Sprintf("/d2l/api/lp/"+c.LPVersion+"/"+format, args...)
}

// lePath builds a path for an LE API route.
func (c *Client) lePath(format string, args ...any) string {
	return fmt.Sprintf("/d2l/api/le/"+c.LEVersion+"/"+format, args...)
}

// basPath builds a path for a BAS API route.
func (c *Client) basPath(format string, args ...any) string {
	return fmt.Sprintf("/d2l/api/bas/"+c.BASVersion+"/"+format, args...)
}

// get executes a GET request, handles rate limiting, stats, and JSON decoding.
func (c *Client) get(path string, params url.Values, out any) error {
	fullURL := c.baseURL + path
	if len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return err
	}

	if err := c.auth.AuthenticateRequest(req); err != nil {
		return err
	}

	// Wait if rate limit is low.
	wait := c.rateLimit.waitIfNeeded()

	start := time.Now()
	resp, err := c.httpClient.Do(req)
	elapsed := time.Since(start)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Update rate limit state from response headers.
	c.rateLimit.update(resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	c.stats.record(elapsed, wait, int64(len(body)))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       string(body),
		}
	}

	if out != nil && len(body) > 0 {
		if err := json.Unmarshal(body, out); err != nil {
			return fmt.Errorf("decoding response: %w", err)
		}
	}
	return nil
}

// getRaw executes a GET request and returns the raw body bytes (for binary downloads).
func (c *Client) getRaw(path string, params url.Values) ([]byte, error) {
	fullURL := c.baseURL + path
	if len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	if err := c.auth.AuthenticateRequest(req); err != nil {
		return nil, err
	}

	wait := c.rateLimit.waitIfNeeded()

	start := time.Now()
	resp, err := c.httpClient.Do(req)
	elapsed := time.Since(start)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	c.rateLimit.update(resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	c.stats.record(elapsed, wait, int64(len(body)))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       string(body),
		}
	}

	return body, nil
}

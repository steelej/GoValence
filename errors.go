package valence

import (
	"fmt"
	"io"
	"net/http"
)

// APIError represents an HTTP error response from the Valence API.
type APIError struct {
	StatusCode int
	Status     string
	Body       string
}

func (e *APIError) Error() string {
	if e.Body != "" {
		return fmt.Sprintf("valence API error %d %s: %s", e.StatusCode, e.Status, e.Body)
	}
	return fmt.Sprintf("valence API error %d %s", e.StatusCode, e.Status)
}

func newAPIError(resp *http.Response) *APIError {
	body, _ := io.ReadAll(resp.Body)
	return &APIError{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       string(body),
	}
}

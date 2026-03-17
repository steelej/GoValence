package valence

import (
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	// creditsPerRequest is the default cost of a single API call.
	creditsPerRequest = 20
	// minCreditsThreshold triggers a sleep before the next request.
	minCreditsThreshold = creditsPerRequest
)

// rateLimitState tracks the current rate-limit window state.
type rateLimitState struct {
	mu        sync.Mutex
	remaining int
	resetAt   time.Time
	initialised bool
}

// update parses rate-limit headers from a response.
func (r *rateLimitState) update(resp *http.Response) {
	remaining := resp.Header.Get("X-Rate-Limit-Remaining")
	reset := resp.Header.Get("X-Rate-Limit-Reset")
	if remaining == "" && reset == "" {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if v, err := strconv.Atoi(remaining); err == nil {
		r.remaining = v
		r.initialised = true
	}
	if v, err := strconv.ParseInt(reset, 10, 64); err == nil {
		r.resetAt = time.Unix(v, 0)
	}
}

// waitIfNeeded sleeps until the rate limit window resets if credits are too low.
// Returns the duration slept (for stats tracking).
func (r *rateLimitState) waitIfNeeded() time.Duration {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.initialised {
		return 0
	}
	if r.remaining > minCreditsThreshold {
		return 0
	}

	sleepUntil := r.resetAt
	if sleepUntil.IsZero() || time.Now().After(sleepUntil) {
		return 0
	}

	wait := time.Until(sleepUntil)
	if wait <= 0 {
		return 0
	}

	r.mu.Unlock()
	time.Sleep(wait)
	r.mu.Lock()

	// Reset our tracked remaining after sleeping.
	r.remaining = 0
	r.initialised = false

	return wait
}

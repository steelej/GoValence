package valence

import (
	"sync"
	"time"
)

// Stats tracks API usage metrics.
type Stats struct {
	mu sync.Mutex

	TotalRequests     int64
	TotalBytes        int64
	TotalDuration     time.Duration
	TotalWaitDuration time.Duration
}

func (s *Stats) record(duration, wait time.Duration, bytes int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.TotalRequests++
	s.TotalBytes += bytes
	s.TotalDuration += duration
	s.TotalWaitDuration += wait
}

// Snapshot returns a copy of the current stats.
func (s *Stats) Snapshot() StatsSnapshot {
	s.mu.Lock()
	defer s.mu.Unlock()
	snap := StatsSnapshot{
		TotalRequests:     s.TotalRequests,
		TotalBytes:        s.TotalBytes,
		TotalDuration:     s.TotalDuration,
		TotalWaitDuration: s.TotalWaitDuration,
	}
	if s.TotalRequests > 0 {
		snap.AvgDuration = s.TotalDuration / time.Duration(s.TotalRequests)
	}
	return snap
}

// Reset zeroes all stats counters.
func (s *Stats) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.TotalRequests = 0
	s.TotalBytes = 0
	s.TotalDuration = 0
	s.TotalWaitDuration = 0
}

// StatsSnapshot is a point-in-time copy of Stats.
type StatsSnapshot struct {
	TotalRequests     int64
	TotalBytes        int64
	TotalDuration     time.Duration
	AvgDuration       time.Duration
	TotalWaitDuration time.Duration
}

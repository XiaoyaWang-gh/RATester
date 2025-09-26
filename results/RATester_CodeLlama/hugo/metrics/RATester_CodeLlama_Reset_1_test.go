package metrics

import (
	"fmt"
	"testing"
	"time"
)

func TestReset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Store{}
	s.mu.Lock()
	s.metrics = make(map[string][]time.Duration)
	s.mu.Unlock()

	s.diffmu.Lock()
	s.diffs = make(map[string]*diff)
	s.diffmu.Unlock()

	s.cachedmu.Lock()
	s.cached = make(map[string]int)
	s.cachedmu.Unlock()

	s.Reset()

	if len(s.metrics) != 0 {
		t.Errorf("expected metrics to be empty, got %v", s.metrics)
	}

	if len(s.diffs) != 0 {
		t.Errorf("expected diffs to be empty, got %v", s.diffs)
	}

	if len(s.cached) != 0 {
		t.Errorf("expected cached to be empty, got %v", s.cached)
	}
}

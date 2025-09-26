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

	s := &Store{
		calculateHints: true,
		metrics:        map[string][]time.Duration{"test": {time.Second}},
		diffs:          map[string]*diff{"test": {}},
		cached:         map[string]int{"test": 1},
	}

	s.Reset()

	if len(s.metrics) != 0 {
		t.Errorf("Expected metrics to be empty after reset, but got %v", s.metrics)
	}

	if len(s.diffs) != 0 {
		t.Errorf("Expected diffs to be empty after reset, but got %v", s.diffs)
	}

	if len(s.cached) != 0 {
		t.Errorf("Expected cached to be empty after reset, but got %v", s.cached)
	}
}

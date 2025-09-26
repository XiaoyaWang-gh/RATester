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
		metrics: map[string][]time.Duration{
			"metric1": {1, 2, 3},
			"metric2": {4, 5, 6},
		},
		diffs: map[string]*diff{
			"diff1": {
				baseline: 10,
				count:    3,
				simSum:   100,
			},
			"diff2": {
				baseline: 20,
				count:    5,
				simSum:   200,
			},
		},
		cached: map[string]int{
			"cached1": 1,
			"cached2": 2,
		},
	}

	s.Reset()

	if len(s.metrics) != 0 {
		t.Errorf("Expected metrics to be empty, got %d", len(s.metrics))
	}

	if len(s.diffs) != 0 {
		t.Errorf("Expected diffs to be empty, got %d", len(s.diffs))
	}

	if len(s.cached) != 0 {
		t.Errorf("Expected cached to be empty, got %d", len(s.cached))
	}
}

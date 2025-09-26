package metrics

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMeasureSince_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &Store{
		calculateHints: false,
		metrics:        make(map[string][]time.Duration),
		mu:             sync.Mutex{},
		diffs:          make(map[string]*diff),
		diffmu:         sync.Mutex{},
		cached:         make(map[string]int),
		cachedmu:       sync.Mutex{},
	}

	key := "testKey"
	start := time.Now()

	store.MeasureSince(key, start)

	store.mu.Lock()
	defer store.mu.Unlock()

	if len(store.metrics[key]) == 0 {
		t.Errorf("Expected metrics for key %s, got none", key)
	}

	duration := store.metrics[key][0]
	if duration < time.Since(start) {
		t.Errorf("Expected duration to be greater than or equal to time.Since(start), got %v", duration)
	}
}

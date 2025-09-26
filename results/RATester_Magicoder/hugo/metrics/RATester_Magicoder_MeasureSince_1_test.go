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

	start := time.Now()
	key := "testKey"
	store.MeasureSince(key, start)

	metrics := store.metrics[key]
	if len(metrics) != 1 {
		t.Errorf("Expected 1 metric, got %d", len(metrics))
	}

	duration := metrics[0]
	if duration < time.Duration(0) {
		t.Errorf("Expected duration to be positive, got %v", duration)
	}
}

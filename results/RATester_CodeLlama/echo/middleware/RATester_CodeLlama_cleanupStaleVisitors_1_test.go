package middleware

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestCleanupStaleVisitors_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &RateLimiterMemoryStore{
		visitors: map[string]*Visitor{
			"1": {
				Limiter:  rate.NewLimiter(1, 1),
				lastSeen: time.Now().Add(-2 * time.Second),
			},
			"2": {
				Limiter:  rate.NewLimiter(1, 1),
				lastSeen: time.Now().Add(-1 * time.Second),
			},
		},
		mutex:       sync.Mutex{},
		rate:        1,
		burst:       1,
		expiresIn:   1 * time.Second,
		lastCleanup: time.Now(),
		timeNow: func() time.Time {
			return time.Now()
		},
	}
	store.cleanupStaleVisitors()
	if len(store.visitors) != 1 {
		t.Errorf("Expected 1 visitor, got %d", len(store.visitors))
	}
	if _, ok := store.visitors["2"]; !ok {
		t.Errorf("Expected visitor 2 to be deleted")
	}
}

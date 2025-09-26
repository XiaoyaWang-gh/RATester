package middleware

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAllow_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &RateLimiterMemoryStore{
		visitors:    make(map[string]*Visitor),
		mutex:       sync.Mutex{},
		rate:        1,
		burst:       1,
		expiresIn:   1 * time.Minute,
		lastCleanup: time.Now(),
		timeNow:     time.Now,
	}

	identifier := "test_identifier"

	allowed, err := store.Allow(identifier)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !allowed {
		t.Errorf("Expected to be allowed, but was not")
	}

	allowed, err = store.Allow(identifier)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if allowed {
		t.Errorf("Expected to be not allowed, but was allowed")
	}

	time.Sleep(2 * time.Minute)

	allowed, err = store.Allow(identifier)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !allowed {
		t.Errorf("Expected to be allowed after cleanup, but was not")
	}
}

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
		rate:        10,
		burst:       20,
		expiresIn:   30 * time.Minute,
		lastCleanup: time.Now(),
		timeNow:     time.Now,
	}

	identifier := "test_identifier"

	allowed, err := store.Allow(identifier)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !allowed {
		t.Errorf("Expected visitor to be allowed, but was not")
	}

	store.visitors[identifier].lastSeen = time.Now().Add(-4 * time.Minute)
	allowed, err = store.Allow(identifier)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !allowed {
		t.Errorf("Expected visitor to be allowed, but was not")
	}

	store.visitors[identifier].lastSeen = time.Now().Add(-31 * time.Minute)
	allowed, err = store.Allow(identifier)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if allowed {
		t.Errorf("Expected visitor to be not allowed, but was allowed")
	}
}

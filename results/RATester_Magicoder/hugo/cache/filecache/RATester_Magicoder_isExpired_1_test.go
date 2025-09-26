package filecache

import (
	"fmt"
	"testing"
	"time"
)

func TestisExpired_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &Cache{
		maxAge: 10 * time.Second,
	}

	modTime := time.Now()

	// Test case 1: maxAge is negative
	cache.maxAge = -1
	if !cache.isExpired(modTime) {
		t.Error("Expected isExpired to return true for maxAge = -1")
	}

	// Test case 2: maxAge is 0
	cache.maxAge = 0
	if !cache.isExpired(modTime) {
		t.Error("Expected isExpired to return true for maxAge = 0")
	}

	// Test case 3: maxAge is greater than time.Since(modTime)
	cache.maxAge = 20 * time.Second
	if cache.isExpired(modTime) {
		t.Error("Expected isExpired to return false for maxAge > 20s")
	}

	// Test case 4: maxAge is less than time.Since(modTime)
	cache.maxAge = 5 * time.Second
	if !cache.isExpired(modTime) {
		t.Error("Expected isExpired to return true for maxAge < 5s")
	}
}

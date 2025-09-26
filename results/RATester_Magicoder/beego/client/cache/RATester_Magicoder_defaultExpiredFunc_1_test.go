package cache

import (
	"fmt"
	"testing"
)

func TestdefaultExpiredFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expiredFunc := defaultExpiredFunc()

	// Test that the function returns a valid duration
	duration := expiredFunc()
	if duration <= 0 {
		t.Errorf("Expected a positive duration, but got %v", duration)
	}

	// Test that the function returns different durations on subsequent calls
	duration1 := expiredFunc()
	if duration1 == duration {
		t.Errorf("Expected different durations, but got the same duration %v", duration)
	}
}

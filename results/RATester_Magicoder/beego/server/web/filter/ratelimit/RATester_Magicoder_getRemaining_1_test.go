package ratelimit

import (
	"fmt"
	"testing"
)

func TestgetRemaining_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &tokenBucket{
		remaining: 10,
	}
	result := b.getRemaining()
	if result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}
}

package ratelimit

import (
	"fmt"
	"testing"
)

func TestgetCapacity_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &tokenBucket{
		capacity: 10,
	}

	capacity := b.getCapacity()
	if capacity != 10 {
		t.Errorf("Expected capacity to be 10, but got %d", capacity)
	}
}

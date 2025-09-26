package ratelimit

import (
	"fmt"
	"testing"
)

func TestGetCapacity_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &tokenBucket{
		capacity: 10,
	}
	if b.getCapacity() != 10 {
		t.Errorf("getCapacity() = %v, want %v", b.getCapacity(), 10)
	}
}

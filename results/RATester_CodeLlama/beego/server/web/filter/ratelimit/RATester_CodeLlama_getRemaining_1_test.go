package ratelimit

import (
	"fmt"
	"testing"
)

func TestGetRemaining_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &tokenBucket{
		remaining: 10,
	}
	if b.getRemaining() != 10 {
		t.Errorf("getRemaining() = %d, want %d", b.getRemaining(), 10)
	}
}

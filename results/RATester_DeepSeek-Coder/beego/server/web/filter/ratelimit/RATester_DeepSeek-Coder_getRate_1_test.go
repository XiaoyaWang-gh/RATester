package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestGetRate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &tokenBucket{
		rate: 10 * time.Second,
	}

	rate := b.getRate()

	if rate != 10*time.Second {
		t.Errorf("Expected rate to be 10 seconds, got %v", rate)
	}
}

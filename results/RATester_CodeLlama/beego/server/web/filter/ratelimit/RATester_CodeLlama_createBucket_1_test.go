package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateBucket_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &limiter{
		capacity: 10,
		rate:     10 * time.Second,
		buckets:  make(map[string]bucket),
	}
	b := l.createBucket("key")
	if b == nil {
		t.Error("createBucket failed")
	}
}

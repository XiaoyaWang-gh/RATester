package ratelimit

import (
	"fmt"
	"testing"
)

func TestWithCapacity_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	capacity := uint(10)
	option := withCapacity(capacity)
	bucket := &tokenBucket{}
	option(bucket)
	if bucket.capacity != capacity {
		t.Errorf("withCapacity failed, want: %d, got: %d", capacity, bucket.capacity)
	}
	if bucket.remaining != capacity {
		t.Errorf("withCapacity failed, want: %d, got: %d", capacity, bucket.remaining)
	}
}

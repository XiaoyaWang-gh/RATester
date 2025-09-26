package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTokenBucket_1(t *testing.T) {
	t.Run("test newTokenBucket", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := newTokenBucket(func(b bucket) {
			b.(*tokenBucket).capacity = 100
			b.(*tokenBucket).rate = time.Second
		})

		if b.getCapacity() != 100 {
			t.Errorf("expected capacity 100, got %d", b.getCapacity())
		}

		if b.getRate() != time.Second {
			t.Errorf("expected rate 1s, got %v", b.getRate())
		}
	})
}

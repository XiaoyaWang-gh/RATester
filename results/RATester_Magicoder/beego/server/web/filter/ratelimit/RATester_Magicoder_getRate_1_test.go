package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestgetRate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &tokenBucket{
		rate: 1 * time.Second,
	}
	got := b.getRate()
	want := 1 * time.Second
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

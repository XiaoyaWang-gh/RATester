package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestRetry_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &ExponentialBackoff{
		InitialInterval: 100 * time.Millisecond,
		MaxBackoffTime:  1 * time.Second,
		Multiplier:      2,
		MaxRetryCount:   3,
	}

	callCount := 0
	err := e.Retry(func() error {
		callCount++
		if callCount < 3 {
			return errors.New("some error")
		}
		return nil
	})

	if err != nil {
		t.Errorf("Retry() returned an error: %v", err)
	}

	if callCount != 3 {
		t.Errorf("Retry() did not retry enough times. Expected 3, got %d", callCount)
	}
}

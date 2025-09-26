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
		MaxBackoffTime:  10 * time.Second,
		Multiplier:      2,
		MaxRetryCount:   5,
	}
	f := func() error {
		return errors.New("error")
	}
	err := e.Retry(f)
	if err == nil {
		t.Errorf("Retry() = %v, want %v", err, "error")
	}
}

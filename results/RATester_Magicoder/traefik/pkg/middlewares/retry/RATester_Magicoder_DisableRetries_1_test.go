package retry

import (
	"fmt"
	"testing"
)

func TestDisableRetries_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := &responseWriter{
		shouldRetry: true,
	}

	rw.DisableRetries()

	if rw.shouldRetry {
		t.Error("Expected shouldRetry to be false after calling DisableRetries, but it was true")
	}
}

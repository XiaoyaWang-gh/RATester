package retry

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDisableRetries_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &responseWriter{
		responseWriter: nil,
		headers:        make(http.Header),
		shouldRetry:    true,
		written:        false,
	}

	r.DisableRetries()

	if r.shouldRetry {
		t.Errorf("Expected shouldRetry to be false, got true")
	}
}

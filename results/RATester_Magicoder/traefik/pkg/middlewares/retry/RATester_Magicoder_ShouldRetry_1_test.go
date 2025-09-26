package retry

import (
	"fmt"
	"testing"
)

func TestShouldRetry_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := &responseWriter{
		shouldRetry: true,
	}

	if !rw.ShouldRetry() {
		t.Error("Expected ShouldRetry to return true, but it did not")
	}

	rw.shouldRetry = false

	if rw.ShouldRetry() {
		t.Error("Expected ShouldRetry to return false, but it did not")
	}
}

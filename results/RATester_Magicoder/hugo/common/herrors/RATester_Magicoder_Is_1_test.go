package herrors

import (
	"fmt"
	"io"
	"testing"
)

func TestIs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &FeatureNotAvailableError{}

	// Test with a different error type
	if e.Is(io.EOF) {
		t.Error("Expected Is to return false for a different error type")
	}

	// Test with the same error type
	if !e.Is(e) {
		t.Error("Expected Is to return true for the same error type")
	}

	// Test with a nil error
	if e.Is(nil) {
		t.Error("Expected Is to return false for a nil error")
	}
}

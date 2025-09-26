package log

import (
	"fmt"
	"testing"
)

func TestDefaultLogger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := DefaultLogger()

	if logger == nil {
		t.Error("Expected logger to be not nil")
	}

	// Add more specific tests for each method in the AllLogger interface
}

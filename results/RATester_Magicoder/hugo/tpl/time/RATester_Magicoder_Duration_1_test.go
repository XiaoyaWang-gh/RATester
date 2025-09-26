package time

import (
	"fmt"
	"testing"
	"time"
)

func TestDuration_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	// Test case 1: Valid input
	duration, err := ns.Duration("s", 10)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if duration != 10*time.Second {
		t.Errorf("Expected duration to be 10 seconds, but got %v", duration)
	}

	// Test case 2: Invalid unit
	_, err = ns.Duration("invalid", 10)
	if err == nil {
		t.Error("Expected an error, but got none")
	}

	// Test case 3: Invalid number
	_, err = ns.Duration("s", "invalid")
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

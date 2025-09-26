package helpers

import (
	"fmt"
	"testing"
)

func TestNewProcessingStats_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "test"
	stats := NewProcessingStats(name)

	if stats.Name != name {
		t.Errorf("Expected name to be %s, but got %s", name, stats.Name)
	}

	if stats.Pages != 0 {
		t.Errorf("Expected Pages to be 0, but got %d", stats.Pages)
	}

	// Add more tests for other fields
}

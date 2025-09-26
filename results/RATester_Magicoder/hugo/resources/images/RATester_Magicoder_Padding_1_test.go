package images

import (
	"fmt"
	"testing"
)

func TestPadding_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filters := &Filters{}

	// Test with valid arguments
	filter := filters.Padding(10, 20, 30, 40, "#ffffff")
	if filter == nil {
		t.Error("Expected filter to be created, but got nil")
	}

	// Test with invalid arguments
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic, but got nil")
		}
	}()
	filters.Padding(10, 20, 30, 40, 50, "#ffffff")
}

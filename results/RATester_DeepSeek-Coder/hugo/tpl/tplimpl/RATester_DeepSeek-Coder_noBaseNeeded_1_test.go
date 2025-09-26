package tplimpl

import (
	"fmt"
	"testing"
)

func TestNoBaseNeeded_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := &templateHandler{}

	tests := []struct {
		name     string
		expected bool
	}{
		{"shortcodes/test", true},
		{"partials/test", true},
		{"test_markup/test", true},
		{"test", false},
	}

	for _, test := range tests {
		result := handler.noBaseNeeded(test.name)
		if result != test.expected {
			t.Errorf("Expected %v, got %v for input %s", test.expected, result, test.name)
		}
	}
}

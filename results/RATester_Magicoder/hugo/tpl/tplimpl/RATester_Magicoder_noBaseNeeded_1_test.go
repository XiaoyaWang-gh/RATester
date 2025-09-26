package tplimpl

import (
	"fmt"
	"testing"
)

func TestnoBaseNeeded_1(t *testing.T) {
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
		{"shortcodes/test.html", true},
		{"partials/test.html", true},
		{"test_markup/test.html", true},
		{"test.html", false},
	}

	for _, test := range tests {
		result := handler.noBaseNeeded(test.name)
		if result != test.expected {
			t.Errorf("Expected %v, but got %v for name %s", test.expected, result, test.name)
		}
	}
}

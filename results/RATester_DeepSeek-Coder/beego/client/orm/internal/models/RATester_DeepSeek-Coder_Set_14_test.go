package models

import (
	"fmt"
	"testing"
)

func TestSet_14(t *testing.T) {
	testCases := []struct {
		name     string
		input    uint32
		expected uint32
	}{
		{"Test case 1", 1, 1},
		{"Test case 2", 12345, 12345},
		{"Test case 3", 4294967295, 4294967295},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			var field PositiveIntegerField
			field.Set(tc.input)
			if field.Value() != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, field.Value())
			}
		})
	}
}

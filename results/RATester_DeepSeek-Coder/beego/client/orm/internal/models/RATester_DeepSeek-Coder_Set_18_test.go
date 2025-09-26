package models

import (
	"fmt"
	"testing"
)

func TestSet_18(t *testing.T) {
	testCases := []struct {
		name     string
		input    int64
		expected int64
	}{
		{"Test case 1", 0, 0},
		{"Test case 2", 1, 1},
		{"Test case 3", -1, -1},
		{"Test case 4", 9223372036854775807, 9223372036854775807},
		{"Test case 5", -9223372036854775808, -9223372036854775808},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			var e BigIntegerField
			e.Set(tc.input)
			if e != BigIntegerField(tc.expected) {
				t.Errorf("Expected %d, got %d", tc.expected, e)
			}
		})
	}
}

package alils

import (
	"fmt"
	"testing"
)

func TestSozLog_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    uint64
		expected int
	}{
		{"Test Case 1", 1, 0},
		{"Test Case 2", 2, 1},
		{"Test Case 3", 3, 1},
		{"Test Case 4", 4, 2},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := sozLog(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}

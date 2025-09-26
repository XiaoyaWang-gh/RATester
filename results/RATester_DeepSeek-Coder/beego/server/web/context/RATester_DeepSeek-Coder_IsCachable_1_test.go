package context

import (
	"fmt"
	"testing"
)

func TestIsCachable_1(t *testing.T) {
	testCases := []struct {
		name     string
		status   int
		expected bool
	}{
		{"200", 200, true},
		{"201", 201, false},
		{"299", 299, true},
		{"300", 300, false},
		{"304", 304, true},
		{"305", 305, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			output := &BeegoOutput{Status: tc.status}
			result := output.IsCachable()
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

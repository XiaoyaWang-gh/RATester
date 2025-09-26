package gin

import (
	"fmt"
	"testing"
)

func TestBodyAllowedForStatus_1(t *testing.T) {
	testCases := []struct {
		name     string
		status   int
		expected bool
	}{
		{"100-199", 199, false},
		{"200-299", 200, true},
		{"300-599", 304, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := bodyAllowedForStatus(tc.status)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

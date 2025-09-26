package try

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestApplyCIMultiplier_1(t *testing.T) {
	os.Setenv("CI", "true")
	defer os.Unsetenv("CI")

	testCases := []struct {
		name     string
		input    time.Duration
		expected time.Duration
	}{
		{
			name:     "Test case 1",
			input:    time.Second,
			expected: time.Duration(float64(time.Second) * CITimeoutMultiplier),
		},
		{
			name:     "Test case 2",
			input:    0,
			expected: 0,
		},
		{
			name:     "Test case 3",
			input:    time.Hour,
			expected: time.Duration(float64(time.Hour) * CITimeoutMultiplier),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := applyCIMultiplier(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

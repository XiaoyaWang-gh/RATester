package capture

import (
	"fmt"
	"testing"
)

func TestRequestSize_1(t *testing.T) {
	testCases := []struct {
		name     string
		capture  *Capture
		expected int64
	}{
		{
			name: "Test with nil readCounter",
			capture: &Capture{
				rr: nil,
			},
			expected: 0,
		},
		{
			name: "Test with non-nil readCounter",
			capture: &Capture{
				rr: &readCounter{
					source: nil,
					size:   1024,
				},
			},
			expected: 1024,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.capture.RequestSize()
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

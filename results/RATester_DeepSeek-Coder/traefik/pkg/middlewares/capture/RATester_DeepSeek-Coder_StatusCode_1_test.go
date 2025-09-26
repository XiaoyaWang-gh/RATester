package capture

import (
	"fmt"
	"testing"
)

func TestStatusCode_1(t *testing.T) {
	testCases := []struct {
		name     string
		status   int
		expected int
	}{
		{
			name:     "StatusCode 200",
			status:   200,
			expected: 200,
		},
		{
			name:     "StatusCode 404",
			status:   404,
			expected: 404,
		},
		{
			name:     "StatusCode 500",
			status:   500,
			expected: 500,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			crw := &captureResponseWriter{
				status: tc.status,
			}
			c := &Capture{
				crw: crw,
			}

			result := c.StatusCode()

			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

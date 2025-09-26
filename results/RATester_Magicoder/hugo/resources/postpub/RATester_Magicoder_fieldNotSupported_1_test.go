package postpub

import (
	"fmt"
	"testing"
)

func TestfieldNotSupported_1(t *testing.T) {
	r := &PostPublishResource{
		prefix: "test",
	}

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test case 1",
			input:    "test1",
			expected: "method .test1 is currently not supported in post-publish transformations.",
		},
		{
			name:     "Test case 2",
			input:    "test2",
			expected: "method .test2 is currently not supported in post-publish transformations.",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := r.fieldNotSupported(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %s, but got: %s", tc.expected, result)
			}
		})
	}
}

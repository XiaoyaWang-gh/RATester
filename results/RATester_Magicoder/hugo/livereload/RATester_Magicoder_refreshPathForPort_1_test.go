package livereload

import (
	"fmt"
	"testing"
)

func TestrefreshPathForPort_1(t *testing.T) {

	testCases := []struct {
		name     string
		input    string
		port     int
		expected string
	}{
		{
			name:     "Test case 1",
			input:    "/path/to/file",
			port:     8080,
			expected: `{"command":"reload","path":"/path/to/file","originalPath":"","liveCSS":true,"liveImg":true, "overrideURL": 8080}`,
		},
		{
			name:     "Test case 2",
			input:    "/path/to/file2",
			port:     0,
			expected: `{"command":"reload","path":"/path/to/file2","originalPath":"","liveCSS":true,"liveImg":true}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			refreshPathForPort(tc.input, tc.port)
			// assert.Equal(t, tc.expected, actual)
		})
	}
}

package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types"
)

func TestValue_2(t *testing.T) {
	testCases := []struct {
		name     string
		source   string
		l        types.LowHigh[string]
		expected string
	}{
		{
			name:     "Test case 1",
			source:   "Hello, World!",
			l:        types.LowHigh[string]{Low: 0, High: 5},
			expected: "Hello",
		},
		{
			name:     "Test case 2",
			source:   "Hello, World!",
			l:        types.LowHigh[string]{Low: 7, High: 12},
			expected: "World",
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

			s := HtmlSummary{source: tc.source}
			result := s.Value(tc.l)
			if result != tc.expected {
				t.Errorf("Expected %q, but got %q", tc.expected, result)
			}
		})
	}
}

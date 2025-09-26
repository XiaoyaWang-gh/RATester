package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/paths"
)

func TestPathInfo_2(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		page     nopPage
		expected *paths.Path
	}

	tests := []testCase{
		{
			name:     "Test 1",
			page:     nopPage(1),
			expected: noOpPathInfo,
		},
		// Add more test cases as needed
	}

	for _, tc := range tests {
		tc := tc // Capture range variable
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			result := tc.page.PathInfo()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

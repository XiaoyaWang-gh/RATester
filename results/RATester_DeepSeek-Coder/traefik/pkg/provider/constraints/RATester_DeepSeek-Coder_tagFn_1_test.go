package constraints

import (
	"fmt"
	"testing"
)

func TestTagFn_1(t *testing.T) {
	testCases := []struct {
		name     string
		tagName  string
		tags     []string
		expected bool
	}{
		{
			name:     "Test case 1",
			tagName:  "tag1",
			tags:     []string{"tag1", "tag2", "tag3"},
			expected: true,
		},
		{
			name:     "Test case 2",
			tagName:  "tag4",
			tags:     []string{"tag1", "tag2", "tag3"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			fn := tagFn(tc.tagName)
			result := fn(tc.tags)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

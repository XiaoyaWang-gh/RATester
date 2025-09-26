package constraints

import (
	"fmt"
	"testing"
)

func TestOrTagFunc_1(t *testing.T) {
	testCases := []struct {
		name     string
		a        constraintTagFunc
		b        constraintTagFunc
		tags     []string
		expected bool
	}{
		{
			name: "both true",
			a: func(tags []string) bool {
				return true
			},
			b: func(tags []string) bool {
				return true
			},
			tags:     []string{"tag1", "tag2"},
			expected: true,
		},
		{
			name: "a true, b false",
			a: func(tags []string) bool {
				return true
			},
			b: func(tags []string) bool {
				return false
			},
			tags:     []string{"tag1", "tag2"},
			expected: true,
		},
		{
			name: "a false, b true",
			a: func(tags []string) bool {
				return false
			},
			b: func(tags []string) bool {
				return true
			},
			tags:     []string{"tag1", "tag2"},
			expected: true,
		},
		{
			name: "both false",
			a: func(tags []string) bool {
				return false
			},
			b: func(tags []string) bool {
				return false
			},
			tags:     []string{"tag1", "tag2"},
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

			result := orTagFunc(tc.a, tc.b)(tc.tags)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

package schema

import (
	"fmt"
	"testing"
)

func TestContains_1(t *testing.T) {
	options := tagOptions{"json", "omitempty"}

	t.Run("Contains", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name     string
			option   string
			expected bool
		}{
			{
				name:     "Option exists",
				option:   "json",
				expected: true,
			},
			{
				name:     "Option does not exist",
				option:   "required",
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

				result := options.Contains(tc.option)
				if result != tc.expected {
					t.Errorf("Expected %v, got %v", tc.expected, result)
				}
			})
		}
	})
}

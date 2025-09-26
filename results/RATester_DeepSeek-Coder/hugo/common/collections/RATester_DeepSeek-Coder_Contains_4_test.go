package collections

import (
	"fmt"
	"testing"
)

func TestContains_4(t *testing.T) {
	slice := SortedStringSlice{"apple", "banana", "cherry", "date", "elderberry"}

	t.Run("Contains", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name     string
			input    string
			expected bool
		}{
			{"Contains apple", "apple", true},
			{"Contains banana", "banana", true},
			{"Contains cherry", "cherry", true},
			{"Contains date", "date", true},
			{"Contains elderberry", "elderberry", true},
			{"Does not contain fig", "fig", false},
			{"Does not contain grape", "grape", false},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				result := slice.Contains(tc.input)
				if result != tc.expected {
					t.Errorf("Expected %v, but got %v", tc.expected, result)
				}
			})
		}
	})
}

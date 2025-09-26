package collections

import (
	"fmt"
	"testing"
)

func TestCount_3(t *testing.T) {
	ss := SortedStringSlice{"apple", "banana", "banana", "cherry", "date", "date", "date"}

	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{"apple", "apple", 1},
		{"banana", "banana", 2},
		{"cherry", "cherry", 1},
		{"date", "date", 3},
		{"empty", "", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ss.Count(tc.s)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

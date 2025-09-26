package pageparser

import (
	"fmt"
	"testing"
)

func TestPos_1(t *testing.T) {
	testCases := []struct {
		name     string
		iterator Iterator
		expected int
	}{
		{
			name: "Test case 1",
			iterator: Iterator{
				items:   Items{},
				lastPos: 0,
			},
			expected: 0,
		},
		{
			name: "Test case 2",
			iterator: Iterator{
				items:   Items{},
				lastPos: 10,
			},
			expected: 10,
		},
		{
			name: "Test case 3",
			iterator: Iterator{
				items:   Items{},
				lastPos: 20,
			},
			expected: 20,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.iterator.Pos()
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}

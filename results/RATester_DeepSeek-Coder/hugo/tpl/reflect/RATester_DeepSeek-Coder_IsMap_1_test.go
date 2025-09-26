package reflect

import (
	"fmt"
	"testing"
)

func TestIsMap_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	testCases := []struct {
		name     string
		input    any
		expected bool
	}{
		{
			name:     "Test with a map",
			input:    map[string]int{"one": 1, "two": 2, "three": 3},
			expected: true,
		},
		{
			name:     "Test with a slice",
			input:    []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "Test with a string",
			input:    "hello",
			expected: false,
		},
		{
			name:     "Test with a nil",
			input:    nil,
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			result := ns.IsMap(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

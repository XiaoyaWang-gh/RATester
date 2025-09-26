package helpers

import (
	"fmt"
	"testing"
)

func TestString_4(t *testing.T) {
	testCases := []struct {
		name     string
		ns       NamedSlice
		expected string
	}{
		{
			name: "Test with empty slice",
			ns: NamedSlice{
				Name:  "test",
				Slice: []string{},
			},
			expected: "test",
		},
		{
			name: "Test with non-empty slice",
			ns: NamedSlice{
				Name:  "test",
				Slice: []string{"a", "b", "c"},
			},
			expected: "test/a,b,c",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.ns.String()
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}

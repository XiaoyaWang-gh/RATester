package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplitNonEscaped_1(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		sep      string
		expected []string
	}{
		{
			name:     "simple split",
			s:        "hello,world",
			sep:      ",",
			expected: []string{"hello", "world"},
		},
		{
			name:     "split with escape",
			s:        "hello\\,world",
			sep:      ",",
			expected: []string{"hello\\,world"},
		},
		{
			name:     "split with multiple escape",
			s:        "hello\\\\,world",
			sep:      ",",
			expected: []string{"hello\\\\", "world"},
		},
		{
			name:     "split with multiple escape and separator",
			s:        "hello\\\\,world,again",
			sep:      ",",
			expected: []string{"hello\\\\", "world", "again"},
		},
		{
			name:     "split with multiple escape and separator 2",
			s:        "hello\\\\\\,world,again",
			sep:      ",",
			expected: []string{"hello\\\\\\,world", "again"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := splitNonEscaped(tc.s, tc.sep)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

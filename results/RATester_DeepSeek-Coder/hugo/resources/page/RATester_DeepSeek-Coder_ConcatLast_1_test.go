package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConcatLast_1(t *testing.T) {
	type testCase struct {
		name     string
		initial  *pagePathBuilder
		input    string
		expected *pagePathBuilder
	}

	testCases := []testCase{
		{
			name: "ConcatLast to empty path",
			initial: &pagePathBuilder{
				els: []string{},
			},
			input: "test",
			expected: &pagePathBuilder{
				els: []string{"test"},
			},
		},
		{
			name: "ConcatLast to non-empty path",
			initial: &pagePathBuilder{
				els: []string{"existing"},
			},
			input: "test",
			expected: &pagePathBuilder{
				els: []string{"existingtest"},
			},
		},
		{
			name: "ConcatLast to path with trailing slash",
			initial: &pagePathBuilder{
				els: []string{"existing/"},
			},
			input: "test",
			expected: &pagePathBuilder{
				els: []string{"existingtest"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tc.initial.ConcatLast(tc.input)
			if !reflect.DeepEqual(tc.initial, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, tc.initial)
			}
		})
	}
}

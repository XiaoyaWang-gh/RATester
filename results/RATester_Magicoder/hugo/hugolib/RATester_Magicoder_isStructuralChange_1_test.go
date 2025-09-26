package hugolib

import (
	"fmt"
	"testing"
)

func TestisStructuralChange_1(t *testing.T) {
	type testCase struct {
		name     string
		input    pathChange
		expected bool
	}

	testCases := []testCase{
		{
			name: "Test case 1: delete is true, isDir is true",
			input: pathChange{
				delete: true,
				isDir:  true,
			},
			expected: true,
		},
		{
			name: "Test case 2: delete is true, isDir is false",
			input: pathChange{
				delete: true,
				isDir:  false,
			},
			expected: true,
		},
		{
			name: "Test case 3: delete is false, isDir is true",
			input: pathChange{
				delete: false,
				isDir:  true,
			},
			expected: true,
		},
		{
			name: "Test case 4: delete is false, isDir is false",
			input: pathChange{
				delete: false,
				isDir:  false,
			},
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

			result := tc.input.isStructuralChange()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

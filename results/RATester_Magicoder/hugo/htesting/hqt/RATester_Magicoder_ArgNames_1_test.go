package hqt

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArgNames_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    argNames
		expected []string
	}{
		{
			name:     "Test case 1",
			input:    argNames{"arg1", "arg2", "arg3"},
			expected: []string{"arg1", "arg2", "arg3"},
		},
		{
			name:     "Test case 2",
			input:    argNames{"arg4", "arg5", "arg6"},
			expected: []string{"arg4", "arg5", "arg6"},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.ArgNames()
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}

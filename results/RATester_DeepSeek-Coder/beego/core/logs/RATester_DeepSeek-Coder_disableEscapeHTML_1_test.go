package logs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDisableEscapeHTML_1(t *testing.T) {
	type testStruct struct {
		escapeHTML bool
	}

	testCases := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name: "Test case 1",
			input: &testStruct{
				escapeHTML: true,
			},
			expected: &testStruct{
				escapeHTML: false,
			},
		},
		{
			name: "Test case 2",
			input: &testStruct{
				escapeHTML: false,
			},
			expected: &testStruct{
				escapeHTML: false,
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

			disableEscapeHTML(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, tc.input)
			}
		})
	}
}

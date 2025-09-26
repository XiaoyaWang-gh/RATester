package pagemeta

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddDateFieldAliases_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Test with single value",
			input:    []string{"field1"},
			expected: []string{"field1", "alias1", "alias2"},
		},
		{
			name:     "Test with multiple values",
			input:    []string{"field1", "field2"},
			expected: []string{"field1", "alias1", "alias2", "field2", "alias3", "alias4"},
		},
		{
			name:     "Test with no matching values",
			input:    []string{"field3"},
			expected: []string{"field3"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			dateFieldAliases = map[string][]string{
				"field1": {"alias1", "alias2"},
				"field2": {"alias3", "alias4"},
			}
			result := addDateFieldAliases(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

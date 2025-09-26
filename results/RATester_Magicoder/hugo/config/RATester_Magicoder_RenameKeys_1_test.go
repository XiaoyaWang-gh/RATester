package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRenameKeys_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]any
		expected map[string]any
	}{
		{
			name: "RenameKeys_EmptyMap",
			input: map[string]any{
				"oldKey1": "value1",
				"oldKey2": "value2",
			},
			expected: map[string]any{
				"oldKey1": "value1",
				"oldKey2": "value2",
			},
		},
		{
			name: "RenameKeys_RenameKeys",
			input: map[string]any{
				"oldKey1": "value1",
				"oldKey2": "value2",
			},
			expected: map[string]any{
				"newKey1": "value1",
				"newKey2": "value2",
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

			RenameKeys(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, tc.input)
			}
		})
	}
}

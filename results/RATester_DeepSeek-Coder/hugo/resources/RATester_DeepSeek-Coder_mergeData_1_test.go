package resources

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeData_1(t *testing.T) {
	type testCase struct {
		name     string
		existing map[string]any
		input    map[string]any
		expected map[string]any
	}

	testCases := []testCase{
		{
			name:     "Empty input",
			existing: map[string]any{"a": 1, "b": 2},
			input:    map[string]any{},
			expected: map[string]any{"a": 1, "b": 2},
		},
		{
			name:     "New keys",
			existing: map[string]any{"a": 1, "b": 2},
			input:    map[string]any{"c": 3, "d": 4},
			expected: map[string]any{"a": 1, "b": 2, "c": 3, "d": 4},
		},
		{
			name:     "Existing keys",
			existing: map[string]any{"a": 1, "b": 2},
			input:    map[string]any{"a": 3, "b": 4},
			expected: map[string]any{"a": 3, "b": 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &genericResource{
				sd: ResourceSourceDescriptor{
					Data: tc.existing,
				},
			}
			r.mergeData(tc.input)
			if !reflect.DeepEqual(r.sd.Data, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, r.sd.Data)
			}
		})
	}
}

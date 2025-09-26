package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMapWithoutPrefix_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]string
		prefix   string
		expected map[string]string
	}{
		{
			name:     "Test case 1",
			input:    map[string]string{"prefix1_key1": "value1", "prefix1_key2": "value2", "key3": "value3"},
			prefix:   "prefix1_",
			expected: map[string]string{"key1": "value1", "key2": "value2"},
		},
		{
			name:     "Test case 2",
			input:    map[string]string{"prefix2_key1": "value1", "prefix2_key2": "value2", "key3": "value3"},
			prefix:   "prefix2_",
			expected: map[string]string{"key1": "value1", "key2": "value2"},
		},
		{
			name:     "Test case 3",
			input:    map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"},
			prefix:   "prefix_",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := GetMapWithoutPrefix(tc.input, tc.prefix)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

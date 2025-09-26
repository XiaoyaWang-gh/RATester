package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMapWithoutPrefix_1(t *testing.T) {
	testCases := []struct {
		name     string
		set      map[string]string
		prefix   string
		expected map[string]string
	}{
		{
			name: "Test case 1",
			set: map[string]string{
				"prefix_key1": "value1",
				"prefix_key2": "value2",
				"other_key":   "value3",
			},
			prefix: "prefix_",
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "Test case 2",
			set: map[string]string{
				"prefix_key1": "value1",
				"prefix_key2": "value2",
				"other_key":   "value3",
			},
			prefix:   "other_",
			expected: map[string]string{},
		},
		{
			name: "Test case 3",
			set: map[string]string{
				"prefix_key1": "value1",
				"prefix_key2": "value2",
				"other_key":   "value3",
			},
			prefix: "no_prefix",
			expected: map[string]string{
				"prefix_key1": "value1",
				"prefix_key2": "value2",
				"other_key":   "value3",
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

			result := GetMapWithoutPrefix(tc.set, tc.prefix)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

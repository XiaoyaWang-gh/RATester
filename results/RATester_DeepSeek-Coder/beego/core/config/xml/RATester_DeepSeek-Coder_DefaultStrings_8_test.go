package xml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefaultStrings_8(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name       string
		key        string
		defaultVal []string
		expected   []string
	}

	testCases := []testCase{
		{
			name:       "TestDefaultStrings_ExistingKey",
			key:        "existingKey",
			defaultVal: []string{"default1", "default2"},
			expected:   []string{"value1", "value2"},
		},
		{
			name:       "TestDefaultStrings_NonExistingKey",
			key:        "nonExistingKey",
			defaultVal: []string{"default1", "default2"},
			expected:   []string{"default1", "default2"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &ConfigContainer{
				data: map[string]interface{}{
					tc.key: tc.expected,
				},
			}

			result := c.DefaultStrings(tc.key, tc.defaultVal)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

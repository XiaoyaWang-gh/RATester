package xml

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestStrings_8(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		key      string
		expected []string
		err      error
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: []string{"value1", "value2"},
			err:      nil,
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: []string{"value3", "value4"},
			err:      nil,
		},
		{
			name:     "Test case 3",
			key:      "key3",
			expected: nil,
			err:      errors.New("key not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			cc := &ConfigContainer{
				data: map[string]interface{}{
					tc.key: tc.expected,
				},
			}

			result, err := cc.Strings(tc.key)
			if !reflect.DeepEqual(result, tc.expected) || !errors.Is(err, tc.err) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}
		})
	}
}

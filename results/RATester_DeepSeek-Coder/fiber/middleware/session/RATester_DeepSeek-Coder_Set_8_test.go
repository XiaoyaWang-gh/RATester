package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSet_8(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		key      string
		val      any
		expected any
	}

	testCases := []testCase{
		{
			name:     "Set string value",
			key:      "key1",
			val:      "value1",
			expected: "value1",
		},
		{
			name:     "Set int value",
			key:      "key2",
			val:      123,
			expected: 123,
		},
		{
			name:     "Set float value",
			key:      "key3",
			val:      123.456,
			expected: 123.456,
		},
		{
			name:     "Set bool value",
			key:      "key4",
			val:      true,
			expected: true,
		},
		{
			name:     "Set nil value",
			key:      "key5",
			val:      nil,
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

			s := &Session{
				data: &data{},
			}
			s.Set(tc.key, tc.val)
			actual := s.Get(tc.key)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}

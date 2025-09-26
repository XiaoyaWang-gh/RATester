package config

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestDIY_4(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected interface{}
		err      error
	}

	testCases := []testCase{
		{
			name:     "test case 1",
			key:      "key1",
			expected: "value1",
			err:      nil,
		},
		{
			name:     "test case 2",
			key:      "key2",
			expected: "value2",
			err:      nil,
		},
		{
			name:     "test case 3",
			key:      "key3",
			expected: nil,
			err:      errors.New("key not find"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &fakeConfigContainer{
				data: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			}
			actual, err := c.DIY(tc.key)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
		})
	}
}

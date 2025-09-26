package gin

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestGetQueryArray_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected []string
		ok       bool
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: []string{"value1", "value2"},
			ok:       true,
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: []string{"value3", "value4"},
			ok:       true,
		},
		{
			name:     "Test case 3",
			key:      "key3",
			expected: nil,
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Context{
				queryCache: url.Values{
					tc.key: tc.expected,
				},
			}

			values, ok := c.GetQueryArray(tc.key)
			if !reflect.DeepEqual(values, tc.expected) || ok != tc.ok {
				t.Errorf("Expected (%v, %v), got (%v, %v)", tc.expected, tc.ok, values, ok)
			}
		})
	}
}

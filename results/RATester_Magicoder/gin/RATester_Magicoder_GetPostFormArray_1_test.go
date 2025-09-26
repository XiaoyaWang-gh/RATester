package gin

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestGetPostFormArray_1(t *testing.T) {
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
			expected: []string{},
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

			ctx := &Context{
				formCache: url.Values{
					tc.key: tc.expected,
				},
			}

			values, ok := ctx.GetPostFormArray(tc.key)

			if !reflect.DeepEqual(values, tc.expected) {
				t.Errorf("Expected values %v, but got %v", tc.expected, values)
			}

			if ok != tc.ok {
				t.Errorf("Expected ok %v, but got %v", tc.ok, ok)
			}
		})
	}
}

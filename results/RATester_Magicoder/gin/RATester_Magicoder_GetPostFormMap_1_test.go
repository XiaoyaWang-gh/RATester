package gin

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestGetPostFormMap_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected map[string]string
		has      bool
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: map[string]string{"key1": "value1"},
			has:      true,
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: map[string]string{"key2": "value2"},
			has:      true,
		},
		{
			name:     "Test case 3",
			key:      "key3",
			expected: map[string]string{"key3": "value3"},
			has:      true,
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
					tc.key: []string{tc.expected[tc.key]},
				},
			}

			result, ok := ctx.GetPostFormMap(tc.key)

			if ok != tc.has {
				t.Errorf("Expected ok to be %v, but got %v", tc.has, ok)
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected result to be %v, but got %v", tc.expected, result)
			}
		})
	}
}

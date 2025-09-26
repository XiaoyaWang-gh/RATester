package gin

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestGetQueryMap_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected map[string]string
		has      bool
		context  *Context
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: map[string]string{"key1": "value1"},
			has:      true,
			context: &Context{
				queryCache: url.Values{"key1": []string{"value1"}},
			},
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: map[string]string{"key2": "value2"},
			has:      true,
			context: &Context{
				queryCache: url.Values{"key2": []string{"value2"}},
			},
		},
		{
			name:     "Test case 3",
			key:      "key3",
			expected: map[string]string{"key3": "value3"},
			has:      true,
			context: &Context{
				queryCache: url.Values{"key3": []string{"value3"}},
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

			result, has := tc.context.GetQueryMap(tc.key)
			if !reflect.DeepEqual(result, tc.expected) || has != tc.has {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

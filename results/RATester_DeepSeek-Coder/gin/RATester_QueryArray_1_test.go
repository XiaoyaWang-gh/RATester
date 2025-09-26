package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestQueryArray_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected []string
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: []string{"value1", "value2"},
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: []string{"value3", "value4"},
		},
		{
			name:     "Test case 3",
			key:      "key3",
			expected: []string{"value5", "value6"},
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
				Request: &http.Request{
					URL: &url.URL{
						RawQuery: fmt.Sprintf("%s=%s", tc.key, strings.Join(tc.expected, ",")),
					},
				},
			}

			result := c.QueryArray(tc.key)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

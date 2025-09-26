package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestQuery_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected string
		context  *Context
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: "value1",
			context: &Context{
				Request: &http.Request{
					URL: &url.URL{
						RawQuery: "key1=value1",
					},
				},
			},
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: "value2",
			context: &Context{
				Request: &http.Request{
					URL: &url.URL{
						RawQuery: "key2=value2",
					},
				},
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

			result := tc.context.Query(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

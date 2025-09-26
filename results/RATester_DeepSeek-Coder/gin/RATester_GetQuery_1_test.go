package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestGetQuery_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected string
		ok       bool
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			key:      "testKey1",
			expected: "testValue1",
			ok:       true,
		},
		{
			name:     "Test case 2",
			key:      "testKey2",
			expected: "testValue2",
			ok:       true,
		},
		{
			name:     "Test case 3",
			key:      "testKey3",
			expected: "",
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
				Request: &http.Request{
					URL: &url.URL{
						RawQuery: fmt.Sprintf("%s=%s", tc.key, tc.expected),
					},
				},
			}

			value, ok := c.GetQuery(tc.key)
			if value != tc.expected || ok != tc.ok {
				t.Errorf("Expected (%s, %t), got (%s, %t)", tc.expected, tc.ok, value, ok)
			}
		})
	}
}

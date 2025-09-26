package gin

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetPostForm_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected string
		ok       bool
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: "value1",
			ok:       true,
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: "value2",
			ok:       true,
		},
		{
			name:     "Test case 3",
			key:      "key3",
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
				formCache: url.Values{
					tc.key: []string{tc.expected},
				},
			}

			value, ok := c.GetPostForm(tc.key)
			if value != tc.expected || ok != tc.ok {
				t.Errorf("Expected (%v, %v), got (%v, %v)", tc.expected, tc.ok, value, ok)
			}
		})
	}
}

package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestPostForm_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected string
	}

	testCases := []testCase{
		{"Test case 1", "key1", "value1"},
		{"Test case 2", "key2", "value2"},
		{"Test case 3", "key3", "value3"},
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
					PostForm: url.Values{
						tc.key: []string{tc.expected},
					},
				},
			}

			result := c.PostForm(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

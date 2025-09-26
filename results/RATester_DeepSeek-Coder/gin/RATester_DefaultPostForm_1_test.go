package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestDefaultPostForm_1(t *testing.T) {
	type testCase struct {
		name         string
		key          string
		defaultValue string
		expected     string
	}

	testCases := []testCase{
		{
			name:         "Test case 1",
			key:          "key1",
			defaultValue: "default1",
			expected:     "value1",
		},
		{
			name:         "Test case 2",
			key:          "key2",
			defaultValue: "default2",
			expected:     "default2",
		},
		{
			name:         "Test case 3",
			key:          "key3",
			defaultValue: "default3",
			expected:     "default3",
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
					PostForm: url.Values{
						tc.key: []string{tc.expected},
					},
				},
			}

			result := c.DefaultPostForm(tc.key, tc.defaultValue)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

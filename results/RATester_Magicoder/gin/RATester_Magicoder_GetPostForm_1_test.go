package gin

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetPostForm_1(t *testing.T) {
	type test struct {
		name     string
		key      string
		expected string
		ok       bool
	}

	tests := []test{
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
			expected: "value3",
			ok:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Context{
				formCache: url.Values{
					tt.key: []string{tt.expected},
				},
			}

			value, ok := c.GetPostForm(tt.key)

			if value != tt.expected {
				t.Errorf("Expected value %s, but got %s", tt.expected, value)
			}

			if ok != tt.ok {
				t.Errorf("Expected ok %t, but got %t", tt.ok, ok)
			}
		})
	}
}

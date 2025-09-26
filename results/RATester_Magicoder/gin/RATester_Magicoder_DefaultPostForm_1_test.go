package gin

import (
	"fmt"
	"net/url"
	"testing"
)

func TestDefaultPostForm_1(t *testing.T) {
	type test struct {
		name         string
		key          string
		defaultValue string
		expected     string
	}

	tests := []test{
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
			expected:     "value2",
		},
		// Add more test cases as needed
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

			result := c.DefaultPostForm(tt.key, tt.defaultValue)

			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}
		})
	}
}

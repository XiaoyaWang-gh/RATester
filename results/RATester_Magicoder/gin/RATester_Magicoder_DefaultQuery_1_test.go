package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestDefaultQuery_1(t *testing.T) {
	type test struct {
		name         string
		context      *Context
		key          string
		defaultValue string
		expected     string
	}

	tests := []test{
		{
			name: "Test case 1",
			context: &Context{
				Request: &http.Request{
					URL: &url.URL{
						RawQuery: "key1=value1&key2=value2",
					},
				},
			},
			key:          "key1",
			defaultValue: "default",
			expected:     "value1",
		},
		{
			name: "Test case 2",
			context: &Context{
				Request: &http.Request{
					URL: &url.URL{
						RawQuery: "key1=value1&key2=value2",
					},
				},
			},
			key:          "key3",
			defaultValue: "default",
			expected:     "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tt.context.DefaultQuery(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}
		})
	}
}

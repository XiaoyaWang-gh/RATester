package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestGetQuery_1(t *testing.T) {
	type test struct {
		name     string
		context  *Context
		key      string
		expected string
		ok       bool
	}

	tests := []test{
		{
			name: "Test case 1",
			context: &Context{
				Request: &http.Request{
					URL: &url.URL{
						RawQuery: "key=value",
					},
				},
			},
			key:      "key",
			expected: "value",
			ok:       true,
		},
		{
			name: "Test case 2",
			context: &Context{
				Request: &http.Request{
					URL: &url.URL{
						RawQuery: "key=value",
					},
				},
			},
			key:      "not_exist",
			expected: "",
			ok:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			value, ok := tt.context.GetQuery(tt.key)
			if value != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, value)
			}
			if ok != tt.ok {
				t.Errorf("Expected %t, but got %t", tt.ok, ok)
			}
		})
	}
}

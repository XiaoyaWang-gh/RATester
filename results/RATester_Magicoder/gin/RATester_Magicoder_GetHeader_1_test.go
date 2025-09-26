package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetHeader_1(t *testing.T) {
	type test struct {
		name     string
		key      string
		expected string
		context  *Context
	}

	tests := []test{
		{
			name:     "TestGetHeader_KeyExists",
			key:      "headerKey",
			expected: "headerValue",
			context: &Context{
				Request: &http.Request{
					Header: http.Header{
						"headerKey": []string{"headerValue"},
					},
				},
			},
		},
		{
			name:     "TestGetHeader_KeyDoesNotExist",
			key:      "nonExistentKey",
			expected: "",
			context: &Context{
				Request: &http.Request{
					Header: http.Header{
						"headerKey": []string{"headerValue"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tt.context.GetHeader(tt.key)
			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}
		})
	}
}

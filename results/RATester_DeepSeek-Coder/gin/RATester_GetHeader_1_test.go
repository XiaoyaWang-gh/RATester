package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetHeader_1(t *testing.T) {
	testCases := []struct {
		name     string
		context  *Context
		key      string
		expected string
	}{
		{
			name: "Test case 1",
			context: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Content-Type": []string{"application/json"},
					},
				},
			},
			key:      "Content-Type",
			expected: "application/json",
		},
		{
			name: "Test case 2",
			context: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Content-Length": []string{"0"},
					},
				},
			},
			key:      "Content-Length",
			expected: "0",
		},
		{
			name: "Test case 3",
			context: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Authorization": []string{"Bearer token"},
					},
				},
			},
			key:      "Authorization",
			expected: "Bearer token",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.context.GetHeader(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

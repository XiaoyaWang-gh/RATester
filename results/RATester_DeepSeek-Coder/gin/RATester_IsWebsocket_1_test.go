package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsWebsocket_1(t *testing.T) {
	testCases := []struct {
		name     string
		context  *Context
		expected bool
	}{
		{
			name: "Test case 1",
			context: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Connection": []string{"Upgrade"},
						"Upgrade":    []string{"Websocket"},
					},
				},
			},
			expected: true,
		},
		{
			name: "Test case 2",
			context: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Connection": []string{"Upgrade"},
						"Upgrade":    []string{"NotWebsocket"},
					},
				},
			},
			expected: false,
		},
		{
			name: "Test case 3",
			context: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Connection": []string{"NotUpgrade"},
						"Upgrade":    []string{"Websocket"},
					},
				},
			},
			expected: false,
		},
		{
			name: "Test case 4",
			context: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Connection": []string{"NotUpgrade"},
						"Upgrade":    []string{"NotWebsocket"},
					},
				},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.context.IsWebsocket()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestErrorHandler_1(t *testing.T) {
	testCases := []struct {
		name     string
		code     string
		handler  http.HandlerFunc
		expected *HttpServer
	}{
		{
			name:    "Test case 1",
			code:    "404",
			handler: func(w http.ResponseWriter, r *http.Request) {},
			expected: &HttpServer{
				Handlers: &ControllerRegister{
					routers: map[string]*Tree{
						"404": {},
					},
				},
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ErrorHandler(tc.code, tc.handler)

			// Check if the result is as expected
			if result.Handlers.routers[tc.code] == nil {
				t.Errorf("Expected ErrorHandler to add a handler for code %s, but it didn't", tc.code)
			}

			// Check if the result is as expected
			if result.Handlers.routers[tc.code].prefix != tc.code {
				t.Errorf("Expected ErrorHandler to add a handler with prefix %s, but got %s", tc.code, result.Handlers.routers[tc.code].prefix)
			}
		})
	}
}

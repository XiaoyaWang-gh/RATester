package api

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetTCPMiddlewares_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		request  *http.Request
		handler  Handler
		expected http.ResponseWriter
	}{
		// TODO: Add test cases
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			test.handler.getTCPMiddlewares(test.expected, test.request)
			// TODO: Add assertions
		})
	}
}

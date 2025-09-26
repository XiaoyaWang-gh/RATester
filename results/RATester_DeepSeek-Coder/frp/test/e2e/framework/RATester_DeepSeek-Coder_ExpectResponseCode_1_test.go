package framework

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/test/e2e/pkg/request"
)

func TestExpectResponseCode_1(t *testing.T) {
	testCases := []struct {
		name     string
		code     int
		response *request.Response
		expected bool
	}{
		{
			name:     "Expect 200",
			code:     200,
			response: &request.Response{Code: 200},
			expected: true,
		},
		{
			name:     "Expect 404",
			code:     404,
			response: &request.Response{Code: 404},
			expected: true,
		},
		{
			name:     "Expect 500",
			code:     500,
			response: &request.Response{Code: 500},
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

			ensureFunc := ExpectResponseCode(tc.code)
			result := ensureFunc(tc.response)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

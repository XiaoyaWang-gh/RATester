package framework

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/test/e2e/pkg/request"
)

func TestExpectError_1(t *testing.T) {
	reqExpect := &RequestExpect{
		req: &request.Request{},
		f:   &Framework{},
	}

	testCases := []struct {
		name     string
		expect   bool
		expected bool
	}{
		{
			name:     "Test case 1",
			expect:   true,
			expected: true,
		},
		{
			name:     "Test case 2",
			expect:   false,
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

			reqExpect.ExpectError(tc.expect)
			if reqExpect.expectError != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, reqExpect.expectError)
			}
		})
	}
}

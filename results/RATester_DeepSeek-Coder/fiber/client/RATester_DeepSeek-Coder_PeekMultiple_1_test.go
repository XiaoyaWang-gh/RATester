package client

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestPeekMultiple_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		header   *Header
		key      string
		expected []string
	}

	testCases := []testCase{
		{
			name: "Test case 1",
			header: &Header{
				RequestHeader: &fasthttp.RequestHeader{},
			},
			key:      "Content-Type",
			expected: []string{"application/json"},
		},
		{
			name: "Test case 2",
			header: &Header{
				RequestHeader: &fasthttp.RequestHeader{},
			},
			key:      "Content-Length",
			expected: []string{"123"},
		},
		{
			name: "Test case 3",
			header: &Header{
				RequestHeader: &fasthttp.RequestHeader{},
			},
			key:      "Authorization",
			expected: []string{"Bearer token"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tc.header.RequestHeader.Set(tc.key, tc.expected[0])
			res := tc.header.PeekMultiple(tc.key)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, res)
			}
		})
	}
}

package httplib

import (
	"fmt"
	"testing"
)

func TestBuildParamBody_1(t *testing.T) {
	type testCase struct {
		name     string
		b        *BeegoHTTPRequest
		expected string
	}

	testCases := []testCase{
		{
			name: "Test with empty params",
			b: &BeegoHTTPRequest{
				params: map[string][]string{},
			},
			expected: "",
		},
		{
			name: "Test with one param",
			b: &BeegoHTTPRequest{
				params: map[string][]string{
					"key1": {"value1"},
				},
			},
			expected: "key1=value1",
		},
		{
			name: "Test with multiple params",
			b: &BeegoHTTPRequest{
				params: map[string][]string{
					"key1": {"value1"},
					"key2": {"value2", "value3"},
				},
			},
			expected: "key1=value1&key2=value2&key2=value3",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.b.buildParamBody()
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

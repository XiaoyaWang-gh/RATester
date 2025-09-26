package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectContainSubstring_1(t *testing.T) {
	gomega.RegisterTestingT(t)

	testCases := []struct {
		name     string
		actual   string
		substr   string
		expected interface{}
	}{
		{
			name:     "ExpectContainSubstring_Success",
			actual:   "Hello, World",
			substr:   "World",
			expected: nil,
		},
		{
			name:     "ExpectContainSubstring_Failure",
			actual:   "Hello, World",
			substr:   "Test",
			expected: "Expected\n    <string>: Hello, World\nto contain substring\n    <string>: Test",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ExpectContainSubstring(tc.actual, tc.substr)
		})
	}
}

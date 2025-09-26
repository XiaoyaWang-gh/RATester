package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectEqualValuesWithOffset_1(t *testing.T) {
	gomega.RegisterTestingT(t)

	testCases := []struct {
		name     string
		offset   int
		actual   interface{}
		expected interface{}
		explain  []interface{}
	}{
		{
			name:     "Test case 1",
			offset:   1,
			actual:   "actual value",
			expected: "expected value",
			explain:  []interface{}{"This is an explaination"},
		},
		{
			name:     "Test case 2",
			offset:   2,
			actual:   123,
			expected: 123,
			explain:  []interface{}{"This is another explaination"},
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

			ExpectEqualValuesWithOffset(tc.offset, tc.actual, tc.expected, tc.explain...)
		})
	}
}

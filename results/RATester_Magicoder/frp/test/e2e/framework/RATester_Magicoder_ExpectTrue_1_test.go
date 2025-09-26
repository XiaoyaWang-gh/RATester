package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectTrue_1(t *testing.T) {
	gomega.RegisterTestingT(t)

	testCases := []struct {
		name     string
		actual   interface{}
		expected bool
	}{
		{
			name:     "Test case 1",
			actual:   true,
			expected: true,
		},
		{
			name:     "Test case 2",
			actual:   false,
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

			ExpectTrue(tc.actual)
		})
	}
}

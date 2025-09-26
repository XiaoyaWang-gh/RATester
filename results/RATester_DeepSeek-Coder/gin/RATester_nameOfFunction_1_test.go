package gin

import (
	"fmt"
	"testing"
)

func TestNameOfFunction_1(t *testing.T) {
	testCases := []struct {
		name     string
		function any
		expected string
	}{
		{
			name:     "Test case 1",
			function: func() {},
			expected: "github.com/username/repo.func1",
		},
		{
			name:     "Test case 2",
			function: func(a int, b string) {},
			expected: "github.com/username/repo.func2",
		},
		{
			name:     "Test case 3",
			function: func(a, b, c int) {},
			expected: "github.com/username/repo.func3",
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

			result := nameOfFunction(tc.function)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}

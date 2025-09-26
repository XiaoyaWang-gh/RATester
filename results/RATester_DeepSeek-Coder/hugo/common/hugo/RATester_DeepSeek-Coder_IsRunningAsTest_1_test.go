package hugo

import (
	"fmt"
	"os"
	"testing"
)

func TestIsRunningAsTest_1(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		expected bool
	}{
		{
			name:     "Test with -test.v flag",
			args:     []string{"-test.v"},
			expected: true,
		},
		{
			name:     "Test with -test flag",
			args:     []string{"-test"},
			expected: true,
		},
		{
			name:     "Test without test flags",
			args:     []string{"-v"},
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

			os.Args = tc.args
			result := IsRunningAsTest()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

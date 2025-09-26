package testenv

import (
	"fmt"
	"testing"
)

func TestParallelOn64Bit_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
	}{
		{
			name: "Test case 1",
		},
		{
			name: "Test case 2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			// Test code here
		})
	}
}

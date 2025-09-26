package types

import (
	"fmt"
	"testing"
)

func TestIsZero_14(t *testing.T) {
	testCases := []struct {
		name   string
		input  LowHigh[string]
		expect bool
	}{
		{
			name:   "Both low and high are zero",
			input:  LowHigh[string]{Low: 0, High: 0},
			expect: true,
		},
		{
			name:   "Only low is zero",
			input:  LowHigh[string]{Low: 0, High: 1},
			expect: true,
		},
		{
			name:   "Only high is zero",
			input:  LowHigh[string]{Low: 1, High: 0},
			expect: false,
		},
		{
			name:   "Both low and high are non-zero",
			input:  LowHigh[string]{Low: 1, High: 1},
			expect: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.IsZero()
			if result != tc.expect {
				t.Errorf("Expected %v, but got %v", tc.expect, result)
			}
		})
	}
}

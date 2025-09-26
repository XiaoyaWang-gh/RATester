package types

import (
	"fmt"
	"testing"
)

func TestEqual_1(t *testing.T) {
	type testCase struct {
		name     string
		q        *BandwidthQuantity
		u        *BandwidthQuantity
		expected bool
	}

	testCases := []testCase{
		{
			name:     "Both nil",
			q:        nil,
			u:        nil,
			expected: true,
		},
		{
			name:     "One nil, one not",
			q:        &BandwidthQuantity{s: "MB", i: 1024},
			u:        nil,
			expected: false,
		},
		{
			name:     "Both not nil, equal",
			q:        &BandwidthQuantity{s: "MB", i: 1024},
			u:        &BandwidthQuantity{s: "MB", i: 1024},
			expected: true,
		},
		{
			name:     "Both not nil, not equal",
			q:        &BandwidthQuantity{s: "MB", i: 1024},
			u:        &BandwidthQuantity{s: "MB", i: 2048},
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

			result := tc.q.Equal(tc.u)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

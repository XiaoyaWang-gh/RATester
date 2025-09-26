package udp

import (
	"fmt"
	"testing"
)

func TestWeightGcd_1(t *testing.T) {
	type testCase struct {
		name     string
		servers  []server
		expected int
	}

	testCases := []testCase{
		{
			name: "GCD of weights",
			servers: []server{
				{weight: 10},
				{weight: 20},
				{weight: 30},
			},
			expected: 10,
		},
		{
			name: "GCD of weights with different values",
			servers: []server{
				{weight: 5},
				{weight: 10},
				{weight: 15},
			},
			expected: 5,
		},
		{
			name: "GCD of weights with one server",
			servers: []server{
				{weight: 100},
			},
			expected: 100,
		},
		{
			name: "GCD of weights with zero",
			servers: []server{
				{weight: 0},
				{weight: 0},
				{weight: 0},
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			lb := &WRRLoadBalancer{
				servers: tc.servers,
			}

			gcd := lb.weightGcd()

			if gcd != tc.expected {
				t.Errorf("Expected GCD of weights to be %d, but got %d", tc.expected, gcd)
			}
		})
	}
}

package tcp

import (
	"fmt"
	"testing"
)

func TestWeightGcd_2(t *testing.T) {
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

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			lb := &WRRLoadBalancer{
				servers: test.servers,
			}

			gcd := lb.weightGcd()

			if gcd != test.expected {
				t.Errorf("Expected GCD of weights to be %d, but got %d", test.expected, gcd)
			}
		})
	}
}

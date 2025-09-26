package udp

import (
	"fmt"
	"testing"
)

func TestWeightGcd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lb := &WRRLoadBalancer{
		servers: []server{
			{weight: 10},
			{weight: 20},
			{weight: 30},
		},
	}
	gcd := lb.weightGcd()
	if gcd != 10 {
		t.Errorf("Expected gcd of 10, got %d", gcd)
	}
}

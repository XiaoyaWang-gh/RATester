package udp

import (
	"fmt"
	"testing"
)

func TestMaxWeight_1(t *testing.T) {
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

	max := lb.maxWeight()
	if max != 30 {
		t.Errorf("Expected max weight to be 30, got %d", max)
	}
}

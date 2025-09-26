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
			{weight: 1},
			{weight: 2},
			{weight: 3},
		},
	}

	max := lb.maxWeight()
	if max != 3 {
		t.Errorf("Expected max weight to be 3, got %d", max)
	}

	lb.servers[1].weight = 5
	max = lb.maxWeight()
	if max != 5 {
		t.Errorf("Expected max weight to be 5, got %d", max)
	}
}

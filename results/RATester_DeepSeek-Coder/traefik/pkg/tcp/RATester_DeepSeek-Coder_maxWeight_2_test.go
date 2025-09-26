package tcp

import (
	"fmt"
	"testing"
)

func TestMaxWeight_2(t *testing.T) {
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
}

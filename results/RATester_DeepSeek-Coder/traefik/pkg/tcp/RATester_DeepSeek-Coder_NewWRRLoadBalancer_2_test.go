package tcp

import (
	"fmt"
	"testing"
)

func TestNewWRRLoadBalancer_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lb := NewWRRLoadBalancer()

	if lb.index != -1 {
		t.Errorf("Expected index to be -1, got %d", lb.index)
	}

	if lb.servers != nil {
		t.Errorf("Expected servers to be nil, got %v", lb.servers)
	}
}

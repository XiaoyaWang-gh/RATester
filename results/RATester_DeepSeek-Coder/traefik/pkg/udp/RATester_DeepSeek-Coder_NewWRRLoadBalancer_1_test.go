package udp

import (
	"fmt"
	"testing"
)

func TestNewWRRLoadBalancer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lb := NewWRRLoadBalancer()

	if lb.index != -1 {
		t.Errorf("Expected index to be -1, got %d", lb.index)
	}
}

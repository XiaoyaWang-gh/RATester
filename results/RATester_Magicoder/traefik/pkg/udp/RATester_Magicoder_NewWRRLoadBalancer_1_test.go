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
	if lb == nil {
		t.Fatal("NewWRRLoadBalancer() returned nil")
	}
	if lb.index != -1 {
		t.Errorf("NewWRRLoadBalancer() index = %v, want -1", lb.index)
	}
}

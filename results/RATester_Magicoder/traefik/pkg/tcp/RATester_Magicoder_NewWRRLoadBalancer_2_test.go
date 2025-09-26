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
	if lb == nil {
		t.Fatal("NewWRRLoadBalancer() should not return nil")
	}
	if lb.index != -1 {
		t.Fatalf("NewWRRLoadBalancer() should initialize index to -1, but got %d", lb.index)
	}
}

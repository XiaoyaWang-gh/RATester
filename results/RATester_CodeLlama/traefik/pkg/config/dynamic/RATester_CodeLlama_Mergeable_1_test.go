package dynamic

import (
	"fmt"
	"testing"
)

func TestMergeable_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &UDPServersLoadBalancer{}
	loadBalancer := &UDPServersLoadBalancer{}

	if !l.Mergeable(loadBalancer) {
		t.Errorf("Expected %v to be mergeable with %v", l, loadBalancer)
	}
}

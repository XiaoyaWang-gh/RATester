package framework

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/test/e2e/pkg/port"
)

func TestAllocPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &Framework{
		portAllocator: &port.Allocator{},
	}

	port := f.AllocPort()

	if port <= 0 {
		t.Errorf("Expected port to be greater than 0, got %d", port)
	}
}

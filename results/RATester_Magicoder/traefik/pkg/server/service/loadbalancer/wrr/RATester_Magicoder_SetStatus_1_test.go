package wrr

import (
	"context"
	"fmt"
	"testing"
)

func TestSetStatus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	b := &Balancer{
		status: make(map[string]struct{}),
	}

	// Test adding a service
	b.SetStatus(ctx, "service1", true)
	if _, ok := b.status["service1"]; !ok {
		t.Error("Expected service1 to be in status map")
	}

	// Test removing a service
	b.SetStatus(ctx, "service1", false)
	if _, ok := b.status["service1"]; ok {
		t.Error("Expected service1 to be removed from status map")
	}

	// Test adding a service that's already in the map
	b.status["service2"] = struct{}{}
	b.SetStatus(ctx, "service2", true)
	if _, ok := b.status["service2"]; !ok {
		t.Error("Expected service2 to still be in status map")
	}

	// Test removing a service that's not in the map
	b.SetStatus(ctx, "service3", false)
	if _, ok := b.status["service3"]; ok {
		t.Error("Expected service3 to not be in status map")
	}
}

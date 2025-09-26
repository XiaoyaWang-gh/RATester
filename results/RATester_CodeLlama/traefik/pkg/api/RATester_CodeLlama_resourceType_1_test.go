package api

import (
	"fmt"
	"testing"
)

func TestResourceType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := tcpMiddlewareRepresentation{}
	m.Type = "test"
	if m.resourceType() != "test" {
		t.Errorf("resourceType() = %v, want %v", m.resourceType(), "test")
	}
}

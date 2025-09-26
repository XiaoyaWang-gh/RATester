package net

import (
	"fmt"
	"testing"
)

func TestNetwork_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ia := &InternalAddr{}
	network := ia.Network()
	if network != "internal" {
		t.Errorf("Expected 'internal', got '%s'", network)
	}
}

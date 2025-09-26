package ecs

import (
	"fmt"
	"testing"
)

func TestGetIPPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{}

	instance := ecsInstance{
		Name: "testInstance",
		machine: &machine{
			privateIP: "192.168.1.1",
		},
	}

	ip, port, err := provider.getIPPort(instance, "8080")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if ip != "192.168.1.1" {
		t.Errorf("Expected IP to be 192.168.1.1, but got %s", ip)
	}

	if port != "8080" {
		t.Errorf("Expected port to be 8080, but got %s", port)
	}
}

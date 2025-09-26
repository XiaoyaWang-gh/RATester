package ecs

import (
	"fmt"
	"testing"
)

func TestGetPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	instance := ecsInstance{
		Name: "test",
	}
	serverPort := "8080"
	result := getPort(instance, serverPort)
	if result != serverPort {
		t.Errorf("getPort() = %v, want %v", result, serverPort)
	}
}

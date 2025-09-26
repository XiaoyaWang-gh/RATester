package ecs

import (
	"fmt"
	"testing"
)

func TestGetServiceName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	instance := ecsInstance{
		Name: "test-instance",
	}
	expected := "test-instance"
	result := getServiceName(instance)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

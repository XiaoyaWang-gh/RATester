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
	actual := getServiceName(instance)

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

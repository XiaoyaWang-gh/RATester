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
	expected := "internal"
	actual := ia.Network()
	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

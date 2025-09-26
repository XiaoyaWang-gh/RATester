package api

import (
	"fmt"
	"testing"
)

func TestName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := tcpMiddlewareRepresentation{}
	m.Name = "test"
	if m.name() != "test" {
		t.Errorf("Expected name to be test, got %v", m.name())
	}
}

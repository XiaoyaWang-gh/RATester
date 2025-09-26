package api

import (
	"fmt"
	"testing"
)

func TestStatus_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := tcpMiddlewareRepresentation{}
	m.Status = "status"
	if m.status() != "status" {
		t.Errorf("status() = %v, want %v", m.status(), "status")
	}
}

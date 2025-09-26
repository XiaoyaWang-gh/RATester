package client

import (
	"fmt"
	"testing"
)

func TestBoundary_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &Request{
		boundary: "testBoundary",
	}

	if req.Boundary() != "testBoundary" {
		t.Errorf("Expected boundary to be 'testBoundary', but got %s", req.Boundary())
	}
}

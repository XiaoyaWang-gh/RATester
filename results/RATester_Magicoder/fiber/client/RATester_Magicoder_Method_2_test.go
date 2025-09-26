package client

import (
	"fmt"
	"testing"
)

func TestMethod_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &Request{
		method: "GET",
	}

	if req.Method() != "GET" {
		t.Errorf("Expected method to be 'GET', but got '%s'", req.Method())
	}
}

package client

import (
	"fmt"
	"testing"
)

func TestClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	if r.Client() != nil {
		t.Errorf("Client() failed")
	}
}

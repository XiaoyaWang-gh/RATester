package request

import (
	"fmt"
	"testing"
)

func TestPort_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	port := 8080
	r.Port(port)
	if r.port != port {
		t.Errorf("Expected port %d, but got %d", port, r.port)
	}
}

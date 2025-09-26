package request

import (
	"fmt"
	"testing"
)

func TestUDP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	r.UDP()
	if r.protocol != "udp" {
		t.Errorf("Expected protocol to be 'udp', got '%s'", r.protocol)
	}
}

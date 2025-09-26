package request

import (
	"fmt"
	"testing"
)

func TestTCP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	r.TCP()
	if r.protocol != "tcp" {
		t.Errorf("Expected protocol to be 'tcp', got '%s'", r.protocol)
	}
}

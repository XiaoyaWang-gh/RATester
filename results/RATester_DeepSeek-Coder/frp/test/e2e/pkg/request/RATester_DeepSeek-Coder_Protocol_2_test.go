package request

import (
	"fmt"
	"testing"
)

func TestProtocol_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	protocol := "http"
	r.Protocol(protocol)
	if r.protocol != protocol {
		t.Errorf("Expected protocol to be %s, but got %s", protocol, r.protocol)
	}
}

package request

import (
	"fmt"
	"testing"
)

func TestHTTPS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	r.HTTPS()
	if r.protocol != "https" {
		t.Errorf("Expected protocol to be 'https', but got '%s'", r.protocol)
	}
}

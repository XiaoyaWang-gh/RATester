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

	r := &Request{
		protocol: "http",
	}
	r.HTTPS()
	if r.protocol != "https" {
		t.Errorf("HTTPS() failed")
	}
}

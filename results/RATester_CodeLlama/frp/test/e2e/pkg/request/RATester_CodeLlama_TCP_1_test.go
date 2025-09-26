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

	r := &Request{
		addr: "127.0.0.1",
		port: 80,
	}
	r.TCP()
	if r.protocol != "tcp" {
		t.Errorf("r.protocol = %s, want %s", r.protocol, "tcp")
	}
}

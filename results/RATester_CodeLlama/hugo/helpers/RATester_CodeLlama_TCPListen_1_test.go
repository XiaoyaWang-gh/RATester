package helpers

import (
	"fmt"
	"testing"
)

func TestTCPListen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l, addr, err := TCPListen()
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	if addr.Network() != "tcp" {
		t.Errorf("addr.Network() = %q; want %q", addr.Network(), "tcp")
	}
	if addr.String() == "" {
		t.Errorf("addr.String() = %q; want non-empty", addr.String())
	}
}

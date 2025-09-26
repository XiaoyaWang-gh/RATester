package request

import (
	"fmt"
	"testing"
)

func TestAddr_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	addr := "127.0.0.1"
	r.Addr(addr)
	if r.addr != addr {
		t.Errorf("Expected addr to be %s, but got %s", addr, r.addr)
	}
}

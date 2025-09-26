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
	r.Addr("127.0.0.1")
	if r.addr != "127.0.0.1" {
		t.Errorf("Addr() failed, got %s, want %s", r.addr, "127.0.0.1")
	}
}

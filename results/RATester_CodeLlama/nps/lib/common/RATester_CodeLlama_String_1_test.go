package common

import (
	"fmt"
	"testing"
)

func TestString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	addr := &Addr{
		Type: 1,
		Host: "127.0.0.1",
		Port: 8080,
	}
	if addr.String() != "127.0.0.1:8080" {
		t.Errorf("addr.String() = %s, want %s", addr.String(), "127.0.0.1:8080")
	}
}

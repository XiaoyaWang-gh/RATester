package streamserver

import (
	"fmt"
	"testing"
)

func TestWithBindAddr_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	addr := "127.0.0.1:8080"
	opt := WithBindAddr(addr)
	s := &Server{}
	opt(s)
	if s.bindAddr != addr {
		t.Errorf("WithBindAddr failed")
	}
}

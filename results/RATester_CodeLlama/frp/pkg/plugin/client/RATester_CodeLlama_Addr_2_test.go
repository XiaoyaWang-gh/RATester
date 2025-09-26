package client

import (
	"fmt"
	"testing"
)

func TestAddr_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Listener{}
	addr := l.Addr()
	if addr == nil {
		t.Errorf("Addr() = nil, want not nil")
	}
}

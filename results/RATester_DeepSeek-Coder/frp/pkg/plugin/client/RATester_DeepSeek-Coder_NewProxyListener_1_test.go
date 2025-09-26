package client

import (
	"fmt"
	"testing"
)

func TestNewProxyListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := NewProxyListener()

	if l == nil {
		t.Errorf("NewProxyListener() = %v, want not nil", l)
	}

	if l.conns == nil {
		t.Errorf("NewProxyListener().conns = %v, want not nil", l.conns)
	}

	if cap(l.conns) != 64 {
		t.Errorf("NewProxyListener().conns = %v, want cap 64", cap(l.conns))
	}
}

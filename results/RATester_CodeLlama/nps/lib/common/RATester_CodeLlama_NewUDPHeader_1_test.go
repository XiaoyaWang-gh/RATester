package common

import (
	"fmt"
	"testing"
)

func TestNewUDPHeader_1(t *testing.T) {
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
	h := NewUDPHeader(0, 0, addr)
	if h.Rsv != 0 {
		t.Errorf("Rsv = %d, want 0", h.Rsv)
	}
	if h.Frag != 0 {
		t.Errorf("Frag = %d, want 0", h.Frag)
	}
	if h.Addr == nil {
		t.Errorf("Addr = nil, want not nil")
	}
	if h.Addr.Type != 1 {
		t.Errorf("Addr.Type = %d, want 1", h.Addr.Type)
	}
	if h.Addr.Host != "127.0.0.1" {
		t.Errorf("Addr.Host = %s, want 127.0.0.1", h.Addr.Host)
	}
	if h.Addr.Port != 8080 {
		t.Errorf("Addr.Port = %d, want 8080", h.Addr.Port)
	}
}

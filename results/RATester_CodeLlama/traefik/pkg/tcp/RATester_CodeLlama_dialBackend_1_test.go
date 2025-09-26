package tcp

import (
	"fmt"
	"testing"
)

func TestDialBackend_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := Proxy{
		address: "127.0.0.1:8080",
	}

	conn, err := p.dialBackend()
	if err != nil {
		t.Errorf("dialBackend() error = %v, want nil", err)
		return
	}

	if conn == nil {
		t.Errorf("dialBackend() conn = nil, want not nil")
	}
}

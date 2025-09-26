package udp

import (
	"fmt"
	"testing"
)

func TestNewProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	address := "127.0.0.1:8080"
	proxy, err := NewProxy(address)
	if err != nil {
		t.Fatalf("NewProxy() = %v, want nil", err)
	}
	if proxy.target != address {
		t.Errorf("NewProxy() = %v, want %v", proxy.target, address)
	}
}

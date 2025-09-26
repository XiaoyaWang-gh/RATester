package proxy

import (
	"fmt"
	"net"
	"testing"
)

func TestNewHttpsListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}

	httpsListener := NewHttpsListener(listener)

	if httpsListener.parentListener != listener {
		t.Errorf("Expected parentListener to be %v, got %v", listener, httpsListener.parentListener)
	}

	if httpsListener.acceptConn == nil {
		t.Errorf("Expected acceptConn to be initialized")
	}

	if err := httpsListener.Close(); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if err := listener.Close(); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

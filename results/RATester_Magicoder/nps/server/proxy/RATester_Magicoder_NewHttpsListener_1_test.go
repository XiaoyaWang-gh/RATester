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

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	httpsListener := NewHttpsListener(listener)

	if httpsListener == nil {
		t.Fatal("NewHttpsListener returned nil")
	}

	if httpsListener.parentListener != listener {
		t.Error("NewHttpsListener did not set the parentListener correctly")
	}

	if httpsListener.acceptConn == nil {
		t.Error("NewHttpsListener did not initialize the acceptConn channel")
	}

	if err := httpsListener.Close(); err != nil {
		t.Error("Close returned an error: ", err)
	}
}

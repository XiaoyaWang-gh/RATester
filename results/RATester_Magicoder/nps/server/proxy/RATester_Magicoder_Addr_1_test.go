package proxy

import (
	"fmt"
	"net"
	"testing"
)

func TestAddr_1(t *testing.T) {
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

	httpsListener := &HttpsListener{
		parentListener: listener,
	}

	addr := httpsListener.Addr()
	if addr.String() != listener.Addr().String() {
		t.Errorf("Expected %s, got %s", listener.Addr().String(), addr.String())
	}
}

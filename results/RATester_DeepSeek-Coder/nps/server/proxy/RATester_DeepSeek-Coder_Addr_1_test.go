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

	expected := listener.Addr()
	actual := httpsListener.Addr()

	if expected.String() != actual.String() {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

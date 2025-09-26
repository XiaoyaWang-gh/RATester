package grace

import (
	"fmt"
	"net"
	"testing"
)

func TestinternalServe_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srv := &Server{
		state: StateRunning,
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	defer ln.Close()

	go func() {
		err := srv.internalServe(ln)
		if err != nil {
			t.Error(err)
		}
	}()

	// Test case logic here
}

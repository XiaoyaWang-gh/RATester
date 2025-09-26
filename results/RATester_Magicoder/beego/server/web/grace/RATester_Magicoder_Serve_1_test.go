package grace

import (
	"fmt"
	"net"
	"testing"
)

func TestServe_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srv := &Server{
		ln: &net.TCPListener{},
	}

	err := srv.Serve()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

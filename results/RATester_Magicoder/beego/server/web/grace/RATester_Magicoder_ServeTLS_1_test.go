package grace

import (
	"fmt"
	"net"
	"os"
	"syscall"
	"testing"
)

func TestServeTLS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srv := &Server{
		isChild: true,
	}

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatal(err)
	}

	err = srv.ServeTLS(ln)
	if err != nil {
		t.Fatal(err)
	}

	process, err := os.FindProcess(os.Getppid())
	if err != nil {
		t.Fatal(err)
	}

	err = process.Signal(syscall.SIGTERM)
	if err != nil {
		t.Fatal(err)
	}

	go srv.handleSignals()

	err = srv.internalServe(ln)
	if err != nil {
		t.Fatal(err)
	}
}

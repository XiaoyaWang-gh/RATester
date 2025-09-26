package udp

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestListen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// CONTEXT_0
	listenConfig := net.ListenConfig{}

	// CONTEXT_1
	network := "tcp"
	address := "127.0.0.1:0"
	timeout := time.Second

	// CONTEXT_2
	l, err := Listen(listenConfig, network, address, timeout)
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	// CONTEXT_3
	network = "tcp"
	address = "127.0.0.1:0"
	timeout = time.Second

	// CONTEXT_4
	l, err = Listen(listenConfig, network, address, timeout)
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	// CONTEXT_5
	network = "tcp"
	address = "127.0.0.1:0"
	timeout = time.Second

	// CONTEXT_6
	l, err = Listen(listenConfig, network, address, timeout)
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	// CONTEXT_7
	network = "tcp"
	address = "127.0.0.1:0"
	timeout = time.Second

	// CONTEXT_8
	l, err = Listen(listenConfig, network, address, timeout)
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	// CONTEXT_9
	if err != nil {
		t.Fatal(err)
	}
}
